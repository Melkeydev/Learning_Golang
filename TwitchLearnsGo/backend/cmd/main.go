package main

import (
  "os"
  "fmt"
  "log"
  "flag"
  "time"
  "context"
  "net/http"
  "database/sql"
  "backend/models"
  _"github.com/lib/pq"
)

type config struct {
  port int 
  env string 
  db struct {
    dsn string
  }
}

type application struct {
  config config
  logger *log.Logger
}

func connectDB(ctx context.Context, cfg config) (*sql.DB, error) {
  db, err := sql.Open("postgres", cfg.db.dsn)
  if err != nil {
    log.Fatal("unable to connect to the database")
    return nil, err
  }

  (&models.DBModel{DB:db}).CreateTable(ctx)

  return db, nil
}

func main() {
  var cfg config
  flag.IntVar(&cfg.port, "port", 4000, "Server for port to listen on") 
  flag.StringVar(&cfg.env, "env", "development", "App Environment")
  flag.StringVar(&cfg.db.dsn, "dsn", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable", "Database connection string")
  flag.Parse()

  logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

  app := &application {
    config: cfg,
    logger: logger ,
  }

  server := http.Server{
    Addr: fmt.Sprintf(":%d", cfg.port),
    Handler: app.routes(),
    IdleTimeout: time.Minute,
    ReadTimeout: 10*time.Second,
    WriteTimeout: 30*time.Second,
  }

  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  db, err := connectDB(ctx, cfg)
  defer db.Close()
  
  err = server.ListenAndServe()

  if err != nil {
    log.Println(err)
  }
}







