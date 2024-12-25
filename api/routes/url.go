package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shreekumar2901/url-shortener/dto"
	"github.com/shreekumar2901/url-shortener/helpers"
	"github.com/shreekumar2901/url-shortener/service"
)

func ListUrls(c *fiber.Ctx) error {
	// Get the corresponding user id
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	userService := service.UserService{}

	userId, err := userService.GetUserIdByUsername(username)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Print(userId)

	service := service.UrlService{}

	response, err := service.ListUrls(userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func DeleteShortByUrl(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	userService := service.UserService{}
	userId, err := userService.GetUserIdByUsername(username)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	url := c.Query("url")

	service := service.UrlService{}

	url = helpers.EnforeHTTP(url)

	if err := service.DeleteShortByUrl(url, userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "record deleted successfully",
	})
}

func ShortenUrl(c *fiber.Ctx) error {
	body := new(dto.UrlShortenRequestDTO)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "can not parse the body",
		})
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	userService := service.UserService{}
	userId, err := userService.GetUserIdByUsername(username)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Check whether shortened URL exists or not
	// TODO: Bind shortened url for particular user
	service := service.UrlService{}

	response, err := service.ShortenUrl(*body, userId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)

}

func ResolveUrl(c *fiber.Ctx) error {
	short := c.Params("short")

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	userService := service.UserService{}
	userId, err := userService.GetUserIdByUsername(username)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	service := service.UrlService{}

	url, err := service.ResolveUrl(short, userId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Redirect(url, 301)

}
