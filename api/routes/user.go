package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shreekumar2901/url-shortener/dto"
	"github.com/shreekumar2901/url-shortener/service"
	"github.com/shreekumar2901/url-shortener/validator"
)

func CreateUser(c *fiber.Ctx) error {
	userDTO := new(dto.UserRequestDTO)

	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse the request body",
		})
	}

	errorMsgs := validator.RegisterUserValidator(userDTO)

	if len(errorMsgs["errors"]) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errorMsgs)
	}

	service := service.UserService{}
	msg, err := service.CreateUser(userDTO)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Some error occured when  creating user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": msg,
	})

}

func GetUserByUserName(c *fiber.Ctx) error {
	username := c.Params("username")

	service := service.UserService{}

	response, err := service.GetUserByUserName(username)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func DeleteUserByUserName(c *fiber.Ctx) error {
	username := c.Params("username")

	service := service.UserService{}

	msg, err := service.DeleteUserByUserName(username)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": msg,
	})
}
