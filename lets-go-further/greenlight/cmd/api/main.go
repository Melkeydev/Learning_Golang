package main

import (
  "fmt"
  "flag"
  "log"
  "net/http"
  "os"
  "time"
  "context"
  "database/sql"
  _"github.com/lib/pq"
  "github.com/amokstakov/greenlight/internal/data"
)

const version = "1.0.0"

type config struct {
  port int
  env string
  db struct {
    dsn string
    maxOpenConns int
    maxIdleConns int
    maxIdleTime string
  }
}

type application struct {
  config config
  logger *log.Logger
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

  flag.Parse()

  // Initialise logger
  logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

  // Declare and connect to the DB
  db, err := openDB(cfg)
  if err != nil {
    logger.Fatal(err)
  }

  defer db.Close()

  logger.Printf("database connection established")

  // Declare instance of the application
  app := &application{
    config: cfg,
    logger: logger,
    models: data.NewModels(db),
  }

  server := &http.Server{
    Addr: fmt.Sprintf(":%d", cfg.port),
    Handler: app.routes(),
    IdleTimeout: time.Minute,
    ReadTimeout: 30 * time.Second,
    WriteTimeout: 30 * time.Second,
  }

  // Start the HTTP Server
  logger.Printf("Starting %s server on %s", cfg.env, server.Addr)
  err = server.ListenAndServe()
  logger.Fatal(err)
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
