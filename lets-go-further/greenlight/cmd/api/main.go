package main

import (
  "fmt"
  "flag"
  "log"
  "net/http"
  "os"
  "time"
)

const version = "1.0.0"

type config struct {
  port int
  env string
}

type application struct {
  config config
  logger *log.Logger
}

func main() {
  var cfg config

  flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
  flag.StringVar(&cfg.env, "env", "dev", "environment")
  flag.Parse()

  // Initialise logger
  logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

  // Declare instance of the application

  app := &application{
    config: cfg,
    logger: logger,
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
  err := server.ListenAndServe()
  logger.Fatal(err)
}
