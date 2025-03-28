package handlers

// these are only for swagger docs and may not be used in the code

// ErrorResponse represents an error response
type ErrorResponse struct {
	Status string `json:"status" example:"error"`
	Error  struct {
		Message string `json:"message" example:"Error message"`
	} `json:"error"`
}

// CreateDatabaseRequest represents the request body for creating a database
type CreateDatabaseRequest struct {
	Name string `json:"name" example:"my_database" binding:"required"`
}

// DatabaseListResponse represents the response for listing databases
type DatabaseListResponse struct {
	Status    string   `json:"status" example:"success"`
	Databases []string `json:"databases"`
}

// CreateDatabaseResponse represents the response for database creation
type CreateDatabaseResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Database created successfully"`
}

// DropDatabaseResponse represents the response for dropping a database
type DropDatabaseResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Database dropped successfully"`
}

// TableListResponse represents the response for listing tables
type TableListResponse struct {
	Status string      `json:"status" example:"success"`
	Tables []TableInfo `json:"tables"`
}

// TableSchemaResponse represents the response for table schema
type TableSchemaResponse struct {
	Status string       `json:"status" example:"success"`
	Schema []ColumnInfo `json:"schema"`
}

// TableDataResponse represents the response for table data
// BUG: wrong swagger docs example
type TableDataResponse struct {
	Status  string                   `json:"status" example:"success"`
	Columns []string                 `json:"columns"`
	Rows    []map[string]interface{} `json:"rows"`
}
