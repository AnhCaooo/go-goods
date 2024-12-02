// AnhCao 2024
package encode

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// EncodeResponse encodes a value of type `T` into JSON and writes it to the HTTP response.
//
// This function serializes the provided value `v` as JSON, sets the HTTP response content type
// to `application/json`, and writes the serialized data to the response body along with the
// specified HTTP status code. If the encoding process fails, an error is returned.
//
// Usage example:
//
//	func GetTodos(w http.ResponseWriter, r *http.Request) {
//	    // Retrieve the list of todos from the database
//	    todos, err := getTodos()
//	    if err != nil {
//	        http.Error(w, err.Error(), http.StatusInternalServerError)
//	        return
//	    }
//
//	    // Write the todos to the HTTP response as JSON
//	    if err := encode.EncodeResponse(w, http.StatusOK, todos); err != nil {
//	        http.Error(w, err.Error(), http.StatusInternalServerError)
//	        return
//	    }
//	}
//
// Type Parameters:
//   - T: The type of the value to be serialized into JSON.
//
// Parameters:
//   - w: The `http.ResponseWriter` used to send the HTTP response.
//   - status: The HTTP status code to set for the response.
//   - v: The value to encode into JSON and include in the response.
//
// Returns:
//   - err: An error if the encoding process fails; nil otherwise.
func EncodeResponse[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

// DecodeRequest decodes the JSON body of an HTTP request into a specified type `T`.
//
// This function reads the request body, parses it as JSON, and unmarshal the data
// into the provided generic type `T`. If the decoding process fails, it returns an
// error detailing the issue.
//
// Usage example:
//
//	func CreateTodo(w http.ResponseWriter, r *http.Request) {
//	    // Decode the request body into a todoStruct
//	    reqBody, err := encode.DecodeRequest[todoStruct](r)
//	    if err != nil {
//	        http.Error(w, err.Error(), http.StatusBadRequest)
//	        return
//	    }
//
//	    // Continue with other business logic
//	}
//
// Type Parameters:
//   - T: The target type into which the JSON data should be unmarshaled.
//
// Parameters:
//   - r: The HTTP request containing the JSON body.
//
// Returns:
//   - v: The decoded value of type `T`.
//   - err: An error if the decoding process fails; nil otherwise.
func DecodeRequest[T any](r *http.Request) (v T, err error) {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

// DecodeResponse decodes the JSON body of an HTTP response into a specified type `T`.
//
// This function reads the response body, parses it as JSON, and unmarshal the data
// into the provided generic type `T`. If the decoding process fails, it returns an
// error describing the issue.
//
// Usage example:
//
//	func FetchTodo() (*todoStruct, error) {
//	    resp, err := http.Get("https://example.com/todo")
//	    if err != nil {
//	        return nil, fmt.Errorf("fetch todo: %w", err)
//	    }
//	    defer resp.Body.Close()
//
//	    // Decode the response body into a todoStruct
//	    todo, err := encode.DecodeResponse[todoStruct](resp)
//	    if err != nil {
//	        return nil, fmt.Errorf("decode response: %w", err)
//	    }
//
//	    return &todo, nil
//	}
//
// Type Parameters:
//   - T: The target type into which the JSON data should be unmarshal.
//
// Parameters:
//   - r: The HTTP response containing the JSON body.
//
// Returns:
//   - v: The decoded value of type `T`.
//   - err: An error if the decoding process fails; nil otherwise.
func DecodeResponse[T any](r *http.Response) (v T, err error) {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}
