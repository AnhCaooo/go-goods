package helpers

import (
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

func TrimSpaceForByte(value []byte) []byte {
	// convert to string
	strVal := string(value)
	trimmedStrVal := strings.TrimSpace(strVal)
	return []byte(trimmedStrVal)
}
