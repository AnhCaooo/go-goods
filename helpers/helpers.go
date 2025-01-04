package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// get current directory (from root to this repo only)
func GetCurrentDir() (dir string, err error) {
	dir, err = os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %s", err.Error())
	}
	return
}

// TrimSpaceForByte parses byte slice to string, then trims all empty space and returns as byte slice.
// This ensures that given text is consistent
func TrimSpaceForByte(value []byte) []byte {
	// convert to string
	strVal := string(value)
	trimmedStrVal := strings.TrimSpace(strVal)
	return []byte(trimmedStrVal)
}

// MapInterfaceToStruct converts an interface{} to a struct of a specified type.
// It first marshals the interface{} to JSON bytes and
// then unmarshals those bytes into the struct.
func MapInterfaceToStruct[T any](data interface{}) (*T, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal interface to JSON: %w", err)
	}

	// Initialize the generic type
	var v T

	// Unmarshal the JSON bytes into the struct
	if err := json.Unmarshal(jsonData, &v); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON to PriceSettings: %w", err)
	}
	return &v, nil
}
