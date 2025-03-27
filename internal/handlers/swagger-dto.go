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
	Databases []string `json:"databases" example:"[\"database1\",\"database2\"]"`
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
