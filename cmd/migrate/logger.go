package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// init zap logger with default fields
func initLogger() *zap.Logger {
	// Setup logger with default fields.
	var zapLoggerConfig zap.Config
	if appEnv == "local" {
		zapLoggerConfig = zap.NewDevelopmentConfig()
	} else {
		zapLoggerConfig = zap.NewProductionConfig()

		zapLoggerConfig.Development = appDebugMode
		zapLoggerConfig.OutputPaths = []string{"stdout"}
		zapLoggerConfig.ErrorOutputPaths = []string{"stdout"}
		zapLoggerConfig.DisableStacktrace = !appDebugMode
		zapLoggerConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	}
	zapLoggerConfig.Level = func() zap.AtomicLevel {
		if appDebugMode {
			return zap.NewAtomicLevelAt(zap.DebugLevel)
		}
		return zap.NewAtomicLevelAt(parseLogLevel(appLogLevel))
	}()
	zapLogger, err := zapLoggerConfig.Build(
		zap.Fields(
			zap.String("app", appName),
			zap.String("build_tag", buildTag),
			zap.String("env", appEnv),
		),
	)
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(zapLogger)

	return zapLogger
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
