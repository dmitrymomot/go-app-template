package main

import (
	"context"
	"fmt"
	stdLog "log"
	"net/http"
	"time"

	"braces.dev/errtrace"
	"github.com/dmitrymomot/asyncer"
	libsql_remote "github.com/dmitrymomot/go-app-template/db/libsql/remote"
	"github.com/dmitrymomot/httpserver"
	"github.com/dmitrymomot/mailer"
	"github.com/dmitrymomot/mailer/adapters/postmark"
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
		mainLogger.Fatalw("Failed to open remote db connection", "error", err)
	}
	defer db.Close()

	// Init redis connection
	redisConnOpt, err := redis.ParseURL(redisConnString)
	if err != nil {
		mainLogger.Fatalw("Failed to parse redis connection string", "error", err)
	}
	redisClient := redis.NewClient(redisConnOpt)
	defer redisClient.Close()

	// Create a new enqueuer with redis as the broker.
	enqueuer := asyncer.MustNewEnqueuer(redisConnString)
	defer enqueuer.Close()

	// Create a new email provider client.
	postmarkAdapter, err := postmark.New(postmarkServerToken, postmarkAccountToken, postmark.Config{
		From:       emailFrom,
		ReplyTo:    emailReplyTo,
		TrackOpens: true,
		TrackLinks: true,
	})
	if err != nil {
		mainLogger.Fatalw("Failed to create postmark adapter", "error", err)
	}
	_ = postmarkAdapter

	// Create a new mail enqueuer.
	mailEnqueuer := mailer.NewEnqueuer(enqueuer)
	_ = mailEnqueuer // TODO: remove this line and use the mailEnqueuer to send emails via the queue.

	// Init router
	r := initRouter(logger, redisClient)

	// TODO: remove this route and add your own instead.
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Hello, World!"))
	})

	eg, ctx := errgroup.WithContext(ctx)

	// Run server
	eg.Go(func() error {
		server := httpserver.New(fmt.Sprintf(":%d", httpPort), r,
			httpserver.WithReadTimeout(httpReadTimeout),
			httpserver.WithWriteTimeout(httpWriteTimeout),
			httpserver.WithGracefulShutdown(10*time.Second),
		)
		return errtrace.Wrap(server.Start(ctx))
	})

	// Run a new queue server with redis as the broker.
	eg.Go(asyncer.RunQueueServer(
		ctx, redisConnString, logger,
		// Register the task handlers.
		mailer.SendEmailHandler(postmarkAdapter), // Register the send_email task handler.
		// ... add more handlers here ...
	))

	// Run a scheduler with redis as the broker.
	// The scheduler will schedule tasks to be enqueued at a specified time.
	eg.Go(asyncer.RunSchedulerServer(
		ctx, redisConnString, logger,
		// Schedule the scheduled_task task to be enqueued every 1 seconds.
		// asyncer.NewTaskScheduler("@every 1s", TestTaskName),
		// ... add more scheduled tasks here ...
	))

	// Wait for all goroutines to finish
	if err := eg.Wait(); err != nil {
		mainLogger.Fatalw("Server stopped with error", "error", err)
	}
}
