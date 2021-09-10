package database

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "github.com/MelkeyDev/models"
)

var DB *gorm.DB

func Connect() {
  dsn := "host=localhost user=postgres password=postgres dbname=melkeydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

  connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  // if there is a error
  if err != nil {
    panic("Could not connect to the DB")
  }

  DB = connection

  connection.AutoMigrate(&models.User{})
}


