package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// Write response to http.ResponseWriter
func RespondJson(w http.ResponseWriter, body any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := ErrorResponse{
			Message: fmt.Sprintf("Failed to get json encoding of body: %v", err),
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Printf("Failed to write error response: %v", err)
		}
		return
	}

	w.WriteHeader(statusCode)
	if _, err := fmt.Fprintf(w, "%s", bodyBytes); err != nil {
		fmt.Printf("Failed to write response: %v", err)
	}
}
