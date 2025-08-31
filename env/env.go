package env

import "os"

// EnvKey represents environment variable key
type EnvKey string

const (
	LogLevel EnvKey = "LOG_LEVEL" // LogLevel represents log level for zap logger
)

// GetValue retrieves the value of environment variable given by key
func (key EnvKey) GetValue() string {
	return os.Getenv(string(key))
}
