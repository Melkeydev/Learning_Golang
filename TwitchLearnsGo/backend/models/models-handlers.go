package models

import (
  "log"
  _"time"
  "context"
  "database/sql"
)

type DBModel struct {
  DB *sql.DB
}

func (db *DBModel) CreateTable(ctx context.Context) error {
  query := `create table if not exists jobs(
    ID INT PRIMARY KEY NOT NULL,
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
