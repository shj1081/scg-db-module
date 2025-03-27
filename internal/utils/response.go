package utils

import (
	"encoding/json"
	"net/http"
)

// send a json response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// send an error response
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]interface{}{
		"status": "error",
		"error": map[string]string{
			"message": message,
		},
	})
}
