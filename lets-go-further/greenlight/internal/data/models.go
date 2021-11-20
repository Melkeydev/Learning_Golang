package data

import (
	"database/sql"
	"errors"
)

// Error for GET method
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// NOTE: This is done differently in UDEMY - please validate

// Models struct that wraps MovieModel
type Models struct {
	Movies      MovieModel
	Users       UserModel
	Tokens      TokenModel
	Permissions PermissionModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies:      MovieModel{DB: db},
		Users:       UserModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Permissions: PermissionModel{DB: db},
	}
}
