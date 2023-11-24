package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/okieLoki/go-fiber/database"
	"github.com/okieLoki/go-fiber/routes"
)

func main() {

	database.ConnectDb()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
