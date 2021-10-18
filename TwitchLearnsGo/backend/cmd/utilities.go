package main

import (
  "log"
  "context"
  "net/http"
  "encoding/json"
  "backend/types"
  "backend/models"
  "database/sql"
)

type envelope map[string]interface{}

func connectDB(ctx context.Context, cfg types.Config) (*sql.DB, error) {
  db, err := sql.Open("postgres", cfg.Db.Dsn)
  if err != nil {
    log.Fatal("unable to connect to the database")
    return nil, err
  }

  (&models.DBModel{DB:db}).CreateJobsTable(ctx)
  (&models.DBModel{DB:db}).CreateUsersTable(ctx)

  return db, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
  js, err := json.MarshalIndent(data, "", "\t")
  if err != nil {
    return err
  }

  js = append(js, '\n')
  
  for k,v := range headers {
    w.Header()[k] = v
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  w.Write(js)
  return nil
}




























