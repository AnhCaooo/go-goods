// AnhCao 2024
package logger

import (
	"time"

	"github.com/AnhCaooo/go-goods/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
initialize and return logger instance (unstructured log). After initialization, use logger as dependency injection.
By default, log level is info. To set the log level, pass environment variable LOG_LEVEL with value DEBUG

*Note*: it is recommended to buffered log entries are flushed before the program exits.

Example usage:

	logger := log.InitLogger()
	defer logger.Sync()
*/
func InitLogger(location *time.Location) *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.In(location).Format("2006-01-02 15:04:05"))
	}

	// Set log level
	level := getLogLevel()
	cfg.Level.SetLevel(zapcore.Level(*level))
	// Disable JSON encoding
	cfg.Encoding = "console"
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

func getLogLevel() *zapcore.Level {
	var logLevel zapcore.Level
	switch env.LogLevel.GetValue() {
	case "DEBUG":
		logLevel = zapcore.DebugLevel
	default:
		logLevel = zapcore.InfoLevel
	}
	return &logLevel
}
