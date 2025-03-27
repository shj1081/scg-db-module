package handlers

// BaseResponse represents the common response structure
type BaseResponse struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message,omitempty" example:"Operation completed successfully"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"Error message"`
}

// CreateDatabaseRequest represents the request body for creating a database
type CreateDatabaseRequest struct {
	Name string `json:"name" example:"my_database" binding:"required"`
}

// DatabaseListData represents the data structure for database list response
type DatabaseListData struct {
	Databases []string `json:"databases" example:"[\"database1\",\"database2\"]"`
}
