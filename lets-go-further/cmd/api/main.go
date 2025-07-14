package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
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
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.databaseUrl, "database-url", os.Getenv("DATABASE_URL"), "Database URL")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	dbpool, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer dbpool.Close()

	logger.Println("database connection poll established")

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on port %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config) (*pgxpool.Pool, error) {
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
	log.Println("Connected to database, current user is:", currentUser)
	return dbPool, nil
}
