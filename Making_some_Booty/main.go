package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
  "github.com/MelkeyDev/routes"
  "github.com/MelkeyDev/database"
)

func main() {
  database.Connect()

  app := fiber.New()

  app.Use(cors.New(cors.Config{
    AllowCredentials: true,
  }))

  // routes
  routes.Test(app)

  app.Listen(":8069")
}

