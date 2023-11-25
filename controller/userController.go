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

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	database.Database.Db.Find(&users)

	if len(users) == 0 {
		return &fiber.Error{
			Code:    fiber.ErrNotFound.Code,
			Message: "No users found",
		}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   users,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User

	id := c.Params("id")

	database.Database.Db.First(&user, id)

	if user.ID == 0 {
		return &fiber.Error{
			Code:    fiber.ErrNotFound.Code,
			Message: "User not found",
		}
	}

	database.Database.Db.Delete(&user)

	return c.Status(200).JSON(fiber.Map{
		"message": "Deleted user",
		"status":  "success",
	})
}
