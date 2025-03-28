package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"scg-inouse-db-module/internal/db"
	"scg-inouse-db-module/internal/utils"

	"github.com/go-chi/chi/v5"
)

// ListTablesHandler godoc
// @Summary List all tables in a database
// @Description Get a list of all tables in the specified database with their information
// @Tags tables
// @Accept json
// @Produce json
// @Param databaseName path string true "Database name"
// @Success 200 {object} TableListResponse "status:success, tables:[]TableInfo"
// @Failure 404 {object} ErrorResponse "Database not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /databases/{databaseName}/tables [get]
func ListTablesHandler(w http.ResponseWriter, r *http.Request) {
	databaseName := chi.URLParam(r, "databaseName")

	// if database not exists, return error
	checkQuery := fmt.Sprintf("SHOW DATABASES LIKE '%s'", databaseName)
	var exists string
	if err := db.DB.QueryRow(checkQuery).Scan(&exists); err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Database not found")
		return
	}

	// get tables info
	query := fmt.Sprintf("SELECT TABLE_NAME, TABLE_ROWS, CREATE_TIME FROM information_schema.tables WHERE TABLE_SCHEMA = '%s'", databaseName)
	rows, err := db.DB.Query(query)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to list tables")
		return
	}
	defer rows.Close()

	// parse tables info
	tables := []TableInfo{}
	for rows.Next() {
		var name sql.NullString
		var rowCount sql.NullInt64
		var createTime sql.NullString

		if err := rows.Scan(&name, &rowCount, &createTime); err != nil {
			log.Printf("Error scanning table info: %v", err)
			utils.RespondWithError(w, http.StatusInternalServerError, "Error scanning table info")
			return
		}
		tables = append(tables, TableInfo{
			Name:      name.String,
			RowCount:  rowCount.Int64,
			CreatedAt: createTime.String,
		})
	}

	// return response
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status": "success",
		"tables": tables,
	})
}

// GetTableSchemaHandler godoc
// @Summary Get table schema
// @Description Get the schema information for a specific table in a database
// @Tags tables
// @Accept json
// @Produce json
// @Param databaseName path string true "Database name"
// @Param tableName path string true "Table name"
// @Success 200 {object} TableSchemaResponse "Schema information"
// @Failure 404 {object} ErrorResponse "Database or table not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /databases/{databaseName}/tables/{tableName}/schema [get]
func GetTableSchemaHandler(w http.ResponseWriter, r *http.Request) {
	databaseName := chi.URLParam(r, "databaseName")
	tableName := chi.URLParam(r, "tableName")

	// if database not exists, return error
	checkQuery := fmt.Sprintf("SHOW DATABASES LIKE '%s'", databaseName)
	var exists string
	if err := db.DB.QueryRow(checkQuery).Scan(&exists); err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Database not found")
		return
	}

	// if table not exists, return error
	checkQuery = fmt.Sprintf("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s'", databaseName, tableName)
	var count int
	if err := db.DB.QueryRow(checkQuery).Scan(&count); err != nil || count == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "Table not found")
		return
	}

	// get table schema
	query := fmt.Sprintf(`
		SELECT 
			c.COLUMN_NAME,
			c.COLUMN_TYPE,
			c.IS_NULLABLE,
			c.COLUMN_DEFAULT,
			c.EXTRA,
			c.COLUMN_KEY,
			kcu.REFERENCED_TABLE_NAME,
			kcu.REFERENCED_COLUMN_NAME
		FROM information_schema.COLUMNS c
		LEFT JOIN information_schema.KEY_COLUMN_USAGE kcu
			ON c.TABLE_SCHEMA = kcu.TABLE_SCHEMA
			AND c.TABLE_NAME = kcu.TABLE_NAME
			AND c.COLUMN_NAME = kcu.COLUMN_NAME
		WHERE c.TABLE_SCHEMA = '%s' 
		AND c.TABLE_NAME = '%s'
		ORDER BY c.ORDINAL_POSITION`, databaseName, tableName)

	rows, err := db.DB.Query(query)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get table schema")
		return
	}
	defer rows.Close()

	// parse table schema
	var columns []ColumnInfo
	for rows.Next() {
		var columnName, columnType, isNullable, columnDefault, columnKey sql.NullString
		var extra string
		var referencedTable, referencedColumn sql.NullString

		if err := rows.Scan(&columnName, &columnType, &isNullable, &columnDefault, &extra,
			&columnKey, &referencedTable, &referencedColumn); err != nil {
			log.Printf("Error scanning schema: %v", err)
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// create column info with required fields
		column := ColumnInfo{
			Name: columnName.String,
			Type: columnType.String,
		}

		// only for exist options
		if isNullable.Valid && isNullable.String != "" {
			column.Nullable = isNullable.String
		}

		if columnDefault.Valid && columnDefault.String != "" {
			column.Default = columnDefault.String
		}

		if extra != "" {
			column.Extra = extra
		}

		if columnKey.Valid && columnKey.String != "" {
			column.Key = columnKey.String
		}

		if referencedTable.Valid && referencedColumn.Valid {
			column.Reference = fmt.Sprintf("%s(%s)", referencedTable.String, referencedColumn.String)
		}

		columns = append(columns, column)
	}

	// return response
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"columns": columns,
	})
}
