package main

import (
	"fiberAuthentication/database"
	"fiberAuthentication/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()
	routes.Setup(app)
	app.Listen(":8000")
}
