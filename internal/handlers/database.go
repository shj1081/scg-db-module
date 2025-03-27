package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"scg-inouse-db-module/internal/db"
	"scg-inouse-db-module/internal/utils"

	"github.com/go-chi/chi/v5"
)

// ListDatabasesHandler godoc
// @Summary List all databases
// @Description Get a list of all available databases
// @Tags databases
// @Accept json
// @Produce json
// @Success 200 {object} BaseResponse{data=DatabaseListData}
// @Failure 500 {object} ErrorResponse
// @Router /databases [get]
func ListDatabasesHandler(w http.ResponseWriter, r *http.Request) {
	var databases []string
	if err := db.DB.Select(&databases, "SHOW DATABASES"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to list databases")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, BaseResponse{
		Status: "success",
		Data: DatabaseListData{
			Databases: databases,
		},
	})
}

// CreateDatabaseHandler godoc
// @Summary Create a new database
// @Description Create a new database with the given name
// @Tags databases
// @Accept json
// @Produce json
// @Param request body CreateDatabaseRequest true "Database creation request"
// @Success 201 {object} BaseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /databases [post]
func CreateDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateDatabaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Malformed JSON")
		return
	}
	query := fmt.Sprintf("CREATE DATABASE %s", req.Name)
	if _, err := db.DB.Exec(query); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create database")
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, BaseResponse{
		Status:  "success",
		Message: "Database created successfully",
	})
}

// DropDatabaseHandler godoc
// @Summary Drop a database
// @Description Drop a database with the given name
// @Tags databases
// @Accept json
// @Produce json
// @Param databaseName path string true "Database name"
// @Success 200 {object} BaseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /databases/{databaseName} [delete]
func DropDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	databaseName := chi.URLParam(r, "databaseName")
	query := fmt.Sprintf("DROP DATABASE %s", databaseName)
	if _, err := db.DB.Exec(query); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to drop database")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, BaseResponse{
		Status:  "success",
		Message: "Database dropped successfully",
	})
}
