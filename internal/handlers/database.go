package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"scg-inouse-db-module/internal/db"
	"scg-inouse-db-module/internal/utils"

	"github.com/go-chi/chi/v5"
)

// list all databases
func ListDatabasesHandler(w http.ResponseWriter, r *http.Request) {
	var databases []string
	if err := db.DB.Select(&databases, "SHOW DATABASES"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to list databases")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":    "success",
		"databases": databases,
	})
}

// create a new database
func CreateDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name string `json:"name"`
	}
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Malformed JSON")
		return
	}
	query := fmt.Sprintf("CREATE DATABASE %s", req.Name)
	if _, err := db.DB.Exec(query); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create database")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Database created successfully",
	})
}

// drop a database
func DropDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	databaseName := chi.URLParam(r, "databaseName")
	query := fmt.Sprintf("DROP DATABASE %s", databaseName)
	if _, err := db.DB.Exec(query); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to drop database")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Database dropped successfully",
	})
}
