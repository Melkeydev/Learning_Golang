package data

import (
	"database/sql"
	"errors"
)

// Error for GET method
var (
	ErrRecordNotFound = errors.New("Record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// NOTE: This is done differently in UDEMY - please validate

// Models struct that wraps MovieModel
type Models struct {
	Movies MovieModel
	Users  UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Users: UserModel{DB: db},
	}
}
