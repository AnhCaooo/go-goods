package encode

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Example struct for testing
type sampleData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func BenchmarkEncodeResponse(b *testing.B) {
	data := sampleData{ID: 1, Name: "Sthor"}

	for b.Loop() {
		// Create a new ResponseRecorder for each iteration to avoid re-use issues
		w := httptest.NewRecorder()

		if err := EncodeResponse(w, http.StatusOK, data); err != nil {
			b.Fatalf("EncodeResponse error: %v", err)
		}
	}
}

func BenchmarkDecodeRequest(b *testing.B) {
	data := sampleData{ID: 1, Name: "Sthor"}
	jsonBytes, _ := json.Marshal(data)

	for b.Loop() {
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonBytes))

		_, err := DecodeRequest[sampleData](r)
		if err != nil {
			b.Fatalf("DecodeRequest error: %v", err)
		}
	}
}
