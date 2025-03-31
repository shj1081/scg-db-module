package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"scg-inouse-db-module/internal/db"
	"scg-inouse-db-module/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/xwb1989/sqlparser"
)

// ensures the identifier (e.g., database name) is alphanumeric with underscores to prevent SQL injection
func isSafeIdentifier(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(s)
}

// walks the AST and adds the database name as a qualifier to all table names
func prefixTableNames(node sqlparser.SQLNode, databaseName string) {
	switch node := node.(type) {
	case *sqlparser.Select:
		for _, tableExpr := range node.From {
			if tbl, ok := tableExpr.(*sqlparser.AliasedTableExpr); ok {
				if tName, ok := tbl.Expr.(sqlparser.TableName); ok && tName.Qualifier.IsEmpty() {
					tName.Qualifier = sqlparser.NewTableIdent(databaseName)
					tbl.Expr = tName
				}
			}
		}
	case *sqlparser.Insert:
		if node.Table.Qualifier.IsEmpty() {
			node.Table.Qualifier = sqlparser.NewTableIdent(databaseName)
		}
	case *sqlparser.Update:
		for _, tableExpr := range node.TableExprs {
			if tbl, ok := tableExpr.(*sqlparser.AliasedTableExpr); ok {
				if tName, ok := tbl.Expr.(sqlparser.TableName); ok && tName.Qualifier.IsEmpty() {
					tName.Qualifier = sqlparser.NewTableIdent(databaseName)
					tbl.Expr = tName
				}
			}
		}
	case *sqlparser.Delete:
		for _, tableExpr := range node.TableExprs {
			if tbl, ok := tableExpr.(*sqlparser.AliasedTableExpr); ok {
				if tName, ok := tbl.Expr.(sqlparser.TableName); ok && tName.Qualifier.IsEmpty() {
					tName.Qualifier = sqlparser.NewTableIdent(databaseName)
					tbl.Expr = tName
				}
			}
		}
	}
}

// handles POST /api/databases/{databaseName}/query
func RawQueryHandler(w http.ResponseWriter, r *http.Request) {
	databaseName := chi.URLParam(r, "databaseName")

	if !isSafeIdentifier(databaseName) {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid database name")
		return
	}

	type request struct {
		Query string `json:"query"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Malformed JSON")
		return
	}

	stmt, err := sqlparser.Parse(req.Query)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid SQL syntax")
		return
	}

	// TODO: Disallow unsupported query types (maybe modified in the future)
	switch stmt.(type) {
	case *sqlparser.DDL:
		utils.RespondWithError(w, http.StatusForbidden, "DDL statements are not allowed")
		return
	}

	prefixTableNames(stmt, databaseName)
	modifiedQuery := sqlparser.String(stmt)

	log.Printf("Executing on DB '%s': %s", databaseName, modifiedQuery)

	switch stmt.(type) {
	case *sqlparser.Select:
		rows, err := db.DB.Queryx(modifiedQuery)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		var results []map[string]interface{}
		for rows.Next() {
			row := make(map[string]interface{})
			if err := rows.MapScan(row); err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, "Row scan failed")
				return
			}
			results = append(results, row)
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
			"status":     "success",
			"resultType": "select",
			"rows":       results,
			"rowCount":   len(results),
		})

	case *sqlparser.Insert, *sqlparser.Update, *sqlparser.Delete:
		res, err := db.DB.Exec(modifiedQuery)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Query execution failed")
			return
		}

		affectedRows, err := res.RowsAffected()
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get affected rows")
			return
		}

		lastInsertId, err := res.LastInsertId()
		if err != nil {
			lastInsertId = 0 // Safe fallback
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
			"status":       "success",
			"resultType":   "update",
			"affectedRows": affectedRows,
			"lastInsertId": lastInsertId,
		})

	default:
		utils.RespondWithError(w, http.StatusBadRequest, "Unsupported query type")
	}
}
