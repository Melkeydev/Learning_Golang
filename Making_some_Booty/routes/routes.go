package routes

import (
  "github.com/gofiber/fiber/v2"
  "github.com/MelkeyDev/controllers"
)

func Test(app *fiber.App) {
  app.Get("/", controllers.Hello)
  app.Post("/api/register", controllers.Regiser)
  app.Post("/api/login", controllers.Login)
}
