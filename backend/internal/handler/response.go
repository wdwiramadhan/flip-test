package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, code string, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := JSONResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
	}
}
