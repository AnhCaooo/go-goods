package goodsHTTP

import (
	"encoding/json"
	"net/http"
)

// HTTPError represents the structure of an HTTP error response.
// This structure is used to provide a consistent error response format across the application.
type HTTPError struct {
	Code           int            `json:"code"`                      // Code is the HTTP status code of the error.
	Error          string         `json:"error"`                     // Error is a short, machine-readable error code. Example: "invalid_request", "unauthorized", "not_found", etc.
	Message        string         `json:"message"`                   // Message is a human-readable description of the error. Example: "The request is invalid", "Unauthorized access", etc. Useful for developer to debug and improve.
	TranslationKey TranslationKey `json:"translation_key,omitempty"` // TranslationKey is an optional field that can be used to provide a key for translation purposes.
}

// Error sends an HTTP error response with the specified status code, message, translation key, and details.
// It sets the Content-Type header to application/json and encodes the HTTPError structure as JSON.
// The status code is set in the response header, and the error message is written to the response
func Error(w http.ResponseWriter, statusCode int, message string, translationKey TranslationKey) {
	httpError := HTTPError{
		Code:    statusCode,
		Error:   http.StatusText(statusCode),
		Message: message,
	}

	if len(translationKey) > 0 {
		httpError.TranslationKey = translationKey
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(httpError); err != nil {
		http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
		return
	}
}
