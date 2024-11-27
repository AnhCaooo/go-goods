// AnhCao 2024
package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
initialize and return logger instance (unstructured log)

*Note*: it is recommended to buffered log entries are flushed before the program exits.

Example usage:

	Logger := log.InitLogger(zapcore.InfoLevel)
	defer Logger.Sync()
*/
func InitLogger(level zapcore.Level) *zap.Logger {
	// todo: do we want to store the log in specific file for investigation purposes?
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = syslogTimeEncoder
	// Set log level
	cfg.Level.SetLevel(zapcore.Level(level))
	// Disable JSON encoding
	cfg.Encoding = "console"

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
