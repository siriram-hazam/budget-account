package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func SendErrorResponse(w http.ResponseWriter, code int, message string, details ...string) {
	w.WriteHeader(code)
	errorDetails := ""
	if len(details) > 0 {
		errorDetails = details[0]
	}
	json.NewEncoder(w).Encode(ErrorResponse{
		Code:    code,
		Message: message,
		Details: errorDetails,
	})
}
