package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dmitrymomot/httpserver"
)

func main() {
	// Init context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init logger with default fields
	log := initLogger()
	defer log.Sync() //nolint:errcheck
	logger := log.Sugar()

	// Init db connection
	db, err := sql.Open("libsql", dbConnString)
	if err != nil {
		logger.Fatalw("Failed to open db connection", "error", err)
	}
	defer db.Close()

	// check db connection
	if err := db.Ping(); err != nil {
		logger.Fatalw("Failed to ping db", "error", err)
	}

	// Init router
	r := initRouter(ctx, db, logger)

	// Run server
	if err := httpserver.Run(ctx, fmt.Sprintf(":%d", httpPort), r); err != nil {
		logger.Fatalw("Failed to start server", "error", err)
	}
}
