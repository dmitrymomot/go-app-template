package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func initRouter(ctx context.Context, db *sql.DB, log *zap.SugaredLogger) http.Handler {
	r := chi.NewRouter()

	// Add middlewares here

	return r
}
