package encode

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Define a sample struct for testing purposes
type Todo struct {
	Title    string `json:"title"`
	Done     bool   `json:"done"`
	Priority int    `json:"priority"`
}

func TestEncodeResponse(t *testing.T) {
	tests := []struct {
		name           string
		status         int         // HTTP status code
		value          interface{} // Value to encode
		expectedBody   string      // Expected JSON body
		expectError    bool        // Whether an error is expected
		expectedErrMsg string      // Expected error message (if any)
	}{
		{
			name:         "Valid JSON",
			status:       http.StatusOK,
			value:        Todo{Title: "Learn Go", Done: false, Priority: 1},
			expectedBody: `{"title":"Learn Go","done":false,"priority":1}` + "\n",
			expectError:  false,
		},
		{
			name:         "Nil Value",
			status:       http.StatusNoContent,
			value:        nil,
			expectedBody: "null\n", // JSON representation of nil
			expectError:  false,
		},
		{
			name:           "Invalid JSON",
			status:         http.StatusInternalServerError,
			value:          make(chan int), // Channels cannot be JSON-encoded
			expectedBody:   "",
			expectError:    true,
			expectedErrMsg: "encode json: json: unsupported type: chan int",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			// Call EncodeResponse
			err := EncodeResponse(rr, test.status, test.value)

			// Check for errors
			if test.expectError {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				if err.Error() != test.expectedErrMsg {
					t.Errorf("expected error message %q, but got %q", test.expectedErrMsg, err.Error())
				}
				return // Skip further checks since the error is expected
			} else {
				if err != nil {
					t.Fatalf("did not expect an error, but got %v", err)
				}
			}

			// Validate response status code
			res := rr.Result()
			if res.StatusCode != test.status {
				t.Errorf("expected status code %d, but got %d", test.status, res.StatusCode)
			}

			// Validate Content-Type header
			if ct := res.Header.Get("Content-Type"); ct != "application/json" {
				t.Errorf("expected Content-Type application/json, but got %q", ct)
			}

			// Validate response body
			body := rr.Body.String()
			if body != test.expectedBody {
				t.Errorf("expected body %q, but got %q", test.expectedBody, body)
			}
		})
	}
}

func TestDecodeRequest(t *testing.T) {
	tests := []struct {
		name       string
		body       interface{}
		expected   Todo
		expectErr  bool
		errMessage string
	}{
		{
			name:      "Valid JSON",
			body:      Todo{Title: "Buy milk", Done: false, Priority: 1},
			expected:  Todo{Title: "Buy milk", Done: false, Priority: 1},
			expectErr: false,
		},
		{
			name:       "Invalid JSON (priority field)",
			body:       map[string]interface{}{"title": "Buy milk", "done": true, "priority": "normal"}, // Invalid JSON
			expected:   Todo{},
			expectErr:  true,
			errMessage: "decode json: json: cannot unmarshal string into Go struct field Todo.priority of type int",
		},
		{
			name:       "Invalid JSON",
			body:       `{"title": "Buy milk", "done": true, "priority"`, // Invalid JSON
			expected:   Todo{},
			expectErr:  true,
			errMessage: "decode json: unexpected EOF",
		},
		{
			name:       "Empty Body",
			body:       nil,
			expected:   Todo{},
			expectErr:  true,
			errMessage: "decode json: EOF",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var requestBody []byte
			if test.body != nil {
				// Serialize the body to JSON if it's not raw
				var err error
				if bodyStr, ok := test.body.(string); ok {
					requestBody = []byte(bodyStr)
				} else {
					requestBody, err = json.Marshal(test.body)
					if err != nil {
						t.Fatalf("failed to marshal test body: %v", err)
					}
				}
			}

			// Create a test HTTP request
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(requestBody))

			// Call DecodeRequest
			decodedResult, err := DecodeRequest[Todo](req)
			// Validate error expectations
			if test.expectErr {
				if err == nil {
					t.Fatalf("expected an error but got nil")
				}
				if err.Error() != test.errMessage {
					t.Errorf("expected error message %q, got %q", test.errMessage, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("did not expect an error but got: %v", err)
				}
				// Validate the result
				if decodedResult != test.expected {
					t.Errorf("expected result %v, got %v", test.expected, decodedResult)
				}
			}
		})
	}
}

func TestDecodeResponse(t *testing.T) {

}
