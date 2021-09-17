package main

import (
  "log"
  "context"
  "backend/types"
  "backend/models"
  "database/sql"
)

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

