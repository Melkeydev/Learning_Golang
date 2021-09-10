package controllers

import (
  "github.com/gofiber/fiber/v2"
  "golang.org/x/crypto/bcrypt"
  "github.com/MelkeyDev/models"
  "github.com/MelkeyDev/database"
)

func Hello(c *fiber.Ctx) error {
  return c.SendString("Midnights Popping")
}

func Regiser(c *fiber.Ctx) error{
  var data map[string]string

  if err := c.BodyParser(&data); err != nil {
    return err
  }

  password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

  user := models.User{
    Name: data["name"],
    Username: data["username"],
    Password: password,
  }

  database.DB.Create(&user)

  return c.JSON(user)
}


func Login(c *fiber.Ctx) error {
  var data map[string]string

  if err := c.BodyParser(&data); err != nil {
    return err
  }

  var user models.User

  database.DB.Where("username = ?", data["username"]).First(&user)

  if user.Id == 0 {
    c.Status(fiber.StatusNotFound)
    return c.JSON(fiber.Map{
      "message" : "User not found",
    })
  }

  if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
    c.Status(fiber.StatusBadRequest)
    return c.JSON(fiber.Map{
      "message": "incorrect credentials",
    })
  }

  return c.JSON(fiber.Map{
    "message" : "success",
  })
}





