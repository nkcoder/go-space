package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"greenlight.danielguo.com/internal/data"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port        int
	env         string
	databaseUrl string
}

type application struct {
	config config
	logger *slog.Logger
	models *data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.databaseUrl, "database-url", os.Getenv("DATABASE_URL"), "Database URL")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	dbpool, err := openDB(cfg, logger)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
	}
	defer dbpool.Close()

	logger.Info("database connection poll established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(dbpool),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Info("Starting server", "env", cfg.env, "addr", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error("Failed to start server", "error", err)
	}
}

func openDB(cfg config, logger *slog.Logger) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), cfg.databaseUrl)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	var currentUser string
	err = dbPool.QueryRow(context.Background(), "SELECT CURRENT_USER").Scan(&currentUser)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to query current user: %v\n", err)
	}
	slog.Info("Connected to database", "current_user", currentUser)
	return dbPool, nil
}
