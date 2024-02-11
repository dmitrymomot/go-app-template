package main

import (
	stdLog "log"

	"github.com/dmitrymomot/go-app-template/pkg/logger"
	"go.uber.org/zap"
)

// init zap logger with default fields
func initLogger() *zap.Logger {
	log, err := logger.New(logger.Config{
		AppEnv:    appEnv,
		Level:     appLogLevel,
		DebugMode: appDebugMode,
		Fields: map[string]interface{}{
			"app":       appName,
			"build_tag": buildTag,
			"env":       appEnv,
		},
	})
	if err != nil {
		stdLog.Fatal(err)
	}
	return log
}
