package main

import (
	"context"
	"database/sql"
	"flag"
	"github.com/amokstakov/greenlight/internal/data"
	"github.com/amokstakov/greenlight/internal/jsonlog"
	_ "github.com/lib/pq"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
	flag.StringVar(&cfg.env, "env", "dev", "environment")
	// This needs to use .env or os
	flag.StringVar(&cfg.db.dsn, "dsn", "host=localhost user=postgres password=postgres dbname=greenlight port=5432 sslmode=disable", "Database connection string")

	// DB configs
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-connections", 25, "postgres max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-connections", 25, "postgres max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "postgres max connection idle time")

	// rate limter settings
	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "rate limiter max requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "rate limiter max burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "enable rate limiter")

	flag.Parse()

	// Initialise logger
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	// Declare and connect to the DB
	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	defer db.Close()

	logger.PrintInfo("database connection established", nil)

	// Declare instance of the application
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// set max number of open connections
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	// set the max number of idle connections
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	// set the max idle timeout
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
