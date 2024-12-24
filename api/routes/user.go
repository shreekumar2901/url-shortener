package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shreekumar2901/url-shortener/config"
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

func Login(c *fiber.Ctx) error {
	loginRquestDTO := new(dto.UserLoginRequestDTO)

	if err := c.BodyParser(&loginRquestDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse body",
		})
	}

	userService := service.UserService{}

	user, err := userService.FindByCredentials(loginRquestDTO.UsernameOrEmail, loginRquestDTO.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized request",
		})
	}

	day := time.Hour * 24

	// Create JWT claims
	claims := jwt.MapClaims{
		"email":    user.Email,
		"username": user.UserName,
		"role":     user.Role,
		"exp":      time.Now().Add(day * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": t,
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
