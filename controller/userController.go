package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/okieLoki/go-fiber/database"
	"github.com/okieLoki/go-fiber/models"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return fiber.ErrBadGateway
	}

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	var existingUser models.User
	database.Database.Db.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.Email != "" {
		return &fiber.Error{
			Code:    fiber.ErrConflict.Code,
			Message: "Email already exists",
		}
	}

	database.Database.Db.Create(&user)

	return c.Status(201).JSON(fiber.Map{
		"message": "Created user",
		"status":  "success",
		"data":    user,
	})
}
