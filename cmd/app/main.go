package main

import (
	"context"
	"fmt"
	stdLog "log"
	"net/http"
	"time"

	"braces.dev/errtrace"
	libsql_remote "github.com/dmitrymomot/go-app-template/db/libsql/remote"
	"github.com/dmitrymomot/httpserver"
	"github.com/redis/go-redis/v9"
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

	// Setup main component logger
	mainLogger := logger.With("component", "main")
	mainLogger.Info("Starting server...")
	defer func() { logger.Info("Server successfully shutdown") }()

	// Init db connection
	db, err := libsql_remote.Connect(dbConnString, dbMaxOpenConns, dbMaxIdleConns)
	if err != nil {
		mainLogger.Fatalw("Failed to open db connection", "error", err)
	}
	defer db.Close()

	// Init redis connection
	redisConnOpt, err := redis.ParseURL(redisConnString)
	if err != nil {
		mainLogger.Fatalw("Failed to parse redis connection string", "error", err)
	}
	redisClient := redis.NewClient(redisConnOpt)
	defer redisClient.Close()

	// Init router
	r := initRouter(logger, redisClient)

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
		return errtrace.Wrap(server.Start(ctx))
	})

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		mainLogger.Fatalw("Server stopped with error", "error", err)
	}
}
