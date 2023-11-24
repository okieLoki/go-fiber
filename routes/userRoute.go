package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/okieLoki/go-fiber/controller"
)

func SetupRoutes(app *fiber.App) {

	route := app.Group("/api/user")

	route.Post("/", controller.CreateUser)
}
