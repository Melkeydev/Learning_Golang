package main

import (
  "os"
  "fmt"
  "log"
  "flag"
  "time"
  "context"
  "net/http"
  "backend/models"
  "backend/types"
  _"github.com/lib/pq"
)

type application struct {
  config types.Config
  logger *log.Logger
  models models.Models
}

func main() {
  var cfg types.Config

  flag.IntVar(&cfg.Port, "port", 4000, "Server for port to listen on") 
  flag.StringVar(&cfg.Env, "env", "development", "App Environment")
  flag.StringVar(&cfg.Db.Dsn, "dsn", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable", "Database connection string")
  flag.Parse()

  logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  db, err := connectDB(ctx, cfg)
  if err != nil {
    log.Println(err)
  }
  defer db.Close()

  app := &application {
    config: cfg,
    logger: logger,
    models: models.NewModels(db),
  }

  server := http.Server{
    Addr: fmt.Sprintf(":%d", cfg.Port),
    Handler: app.routes(),
    IdleTimeout: time.Minute,
    ReadTimeout: 10*time.Second,
    WriteTimeout: 30*time.Second,
  }

  err = server.ListenAndServe()

  if err != nil {
    log.Println(err)
  }
}

