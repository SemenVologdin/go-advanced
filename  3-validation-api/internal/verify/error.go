package verify

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error     bool   `json:"error"`
	ErrorText string `json:"error_text"`
}

func errorResponse(w http.ResponseWriter, text string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(text); err != nil {
		log.Printf("json.NewEncoder.Encode: %v", err)
	}
}
