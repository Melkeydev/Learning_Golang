package data

import (
  "database/sql"
  "errors"
)

// Error for GET method
var (
  ErrRecordNotFound = errors.New("Record not found")
)

// NOTE: This is done differently in UDEMY - please validate

// Models struct that wraps MovieModel
type Models struct {
  Movies MovieModel
}

func NewModels(db *sql.DB) Models {
  return Models {
    Movies: MovieModel{DB: db},
  }
}
