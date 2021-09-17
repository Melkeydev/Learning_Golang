package models

import (
  "log"
  _"time"
  "context"
)


func (db *DBModel) CreateJobsTable(ctx context.Context) error {
  query := `create table if not exists jobs(
    ID SERIAL PRIMARY KEY NOT NULL,
    TITLE TEXT NOT NULL,
    COMPANY TEXT NOT NULL,
    LINK TEXT NOT NULL,
    DESCRIPTION TEXT NOT NULL,
    TOTAL_COMPENSATION INT 
  )`

  _, err := db.DB.ExecContext(ctx, query) 

  if err != nil {
    return err 
  }

  log.Println("Created table")
  return nil
}

func (db *DBModel) CreateUsersTable(ctx context.Context) error {
  query := `create table if not exists users(
    ID SERIAL PRIMARY KEY NOT NULL,
    USERNAME TEXT NOT NULL,
    PASSWORD TEXT NOT NULL
  )`

  _, err := db.DB.ExecContext(ctx, query) 

  if err != nil {
    return err 
  }

  log.Println("Created table")
  return nil
}
























