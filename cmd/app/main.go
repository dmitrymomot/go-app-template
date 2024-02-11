package main

import (
	"context"
	"fmt"
	stdLog "log"
	"net/http"
	"time"

	"github.com/dmitrymomot/go-app-template/db/libsql"
	"github.com/dmitrymomot/go-app-template/db/repository"
	"github.com/dmitrymomot/httpserver"
	"golang.org/x/sync/errgroup"
)

func main() {
	// Init context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init logger with default fields
	log := initLogger()
	defer func() {
		if err := log.Sync(); err != nil {
			stdLog.Printf("Failed to flush log buffer: %v", err)
		}
	}()
	logger := log.Sugar()

	// Init db connection
	db, err := libsql.Connect(dbConnString, dbMaxOpenConns, dbMaxIdleConns)
	if err != nil {
		logger.Fatalw("Failed to open db connection", "error", err)
	}
	defer db.Close()

	repo := repository.New(db)
	_ = repo

	// Init router
	r := initRouter(ctx, db, logger)

	// TODO: remove this route and add your own instead.
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Hello, World!"))
	})

	g, ctx := errgroup.WithContext(ctx)

	// Run server
	g.Go(func() error {
		server := httpserver.New(fmt.Sprintf(":%d", httpPort), r,
			httpserver.WithReadTimeout(httpReadTimeout),
			httpserver.WithWriteTimeout(httpWriteTimeout),
			httpserver.WithGracefulShutdown(10*time.Second),
		)
		return server.Start(ctx)
	})

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		logger.Fatalw("Server stopped with error", "error", err)
	}
}
