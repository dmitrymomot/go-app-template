package logger

import (
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config defines the configuration for the logger.
type Config struct {
	AppEnv    string
	Level     string
	DebugMode bool
	Global    bool
	Fields    map[string]interface{}
}

// New creates a new logger instance based on the provided configuration.
// It returns a pointer to a zap.Logger and an error if the logger initialization fails.
// The logger is configured based on the provided Config struct, which includes options such as
// the application environment, debug mode, log level, additional fields, and whether to replace the global logger.
func New(cnf Config) (*zap.Logger, error) {
	// Setup logger with default fields.
	var zapLoggerConfig zap.Config
	if cnf.AppEnv == "local" || cnf.AppEnv == "development" {
		zapLoggerConfig = zap.NewDevelopmentConfig()
	} else {
		zapLoggerConfig = zap.NewProductionConfig()

		zapLoggerConfig.Development = cnf.DebugMode
		zapLoggerConfig.OutputPaths = []string{"stdout"}
		zapLoggerConfig.ErrorOutputPaths = []string{"stdout"}
		zapLoggerConfig.DisableStacktrace = !cnf.DebugMode
		zapLoggerConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	}
	zapLoggerConfig.Level = func() zap.AtomicLevel {
		if cnf.DebugMode {
			return zap.NewAtomicLevelAt(zap.DebugLevel)
		}
		return zap.NewAtomicLevelAt(parseLogLevel(cnf.Level))
	}()

	opts := make([]zap.Option, 0)

	// Set up the fields.
	if fieldsCount := len(cnf.Fields); fieldsCount > 0 {
		fields := make([]zap.Field, 0, fieldsCount)
		for k, v := range cnf.Fields {
			fields = append(fields, zap.Any(k, v))
		}
		opts = append(opts, zap.Fields(fields...))
	}

	// Build the logger.
	zapLogger, err := zapLoggerConfig.Build(opts...)
	if err != nil {
		return nil, errors.Join(ErrFailedToInitLogger, err)
	}

	// Replace the global logger.
	if cnf.Global {
		zap.ReplaceGlobals(zapLogger)
	}

	return zapLogger, nil
}

// parseLogLevel parses the log level from the string.
func parseLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	default:
		return zap.ErrorLevel
	}
}
