package models

import (
  //_"time"
  "database/sql"
)

type Models struct {
  DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type DBModel struct {
  DB *sql.DB
}

type Job struct {
  ID int `json:"id"`
  Title string `json:"title"`
  Company string `json:"company"`
  Link string `json:"link"`
  Description string `json:"description"`
  TotalCompensation int `json:"total_compensation"`
}

type User struct {
  ID int `json:"id"`
  Username string `json:"username"`
  Password string `json:"-"`
}






































