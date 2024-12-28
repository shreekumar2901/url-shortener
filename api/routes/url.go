package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shreekumar2901/url-shortener/dto"
	"github.com/shreekumar2901/url-shortener/helpers"
	"github.com/shreekumar2901/url-shortener/service"
)

// @Summary List Urls ans shorts for the User
// @Description List the all urls and their short for the user
// @Tags Url
// @Produce json
// @Param Authorization header string true "Bearer Token in the format 'Bearer <token>'"
// @Success 200 {array} []dto.UrlListResponseDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/urls [get]
func ListUrls(c *fiber.Ctx) error {
	// Get the corresponding user id
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	userService := service.UserService{}

	userId, err := userService.GetUserIdByUsername(username)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			ErrorMsgs:  []string{err.Error()},
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	service := service.UrlService{}

	response, err := service.ListUrls(userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			ErrorMsgs:  []string{err.Error()},
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary Delete Short By Url
// @Description Delete the short for given url
// @Tags Url
// @Produce json
// @Param Authorization header string true "Bearer Token in the format 'Bearer <token>'"
// @Param url query string true "The URL for which the short URL should be deleted"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/urls [delete]
func DeleteShortByUrl(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	userService := service.UserService{}
	userId, err := userService.GetUserIdByUsername(username)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			ErrorMsgs:  []string{err.Error()},
			StatusCode: fiber.StatusBadRequest,
		})
	}

	url := c.Query("url")

	service := service.UrlService{}

	url = helpers.EnforeHTTP(url)

	if err := service.DeleteShortByUrl(url, userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			ErrorMsgs:  []string{err.Error()},
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "record deleted successfully",
	})
}

// @Summary Create Short for a URL
// @Description Creates a custom short for given URL
// @Tags Url
// @Produce json
// @Param Authorization header string true "Bearer Token in the format 'Bearer <token>'"
// @Param request body dto.UrlShortenRequestDTO true "Request body for creating a short URL"
// @Success 200 {object} dto.UrlShortenResponseDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/urls/shorten [post]
func ShortenUrl(c *fiber.Ctx) error {
	body := new(dto.UrlShortenRequestDTO)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			ErrorMsgs:  []string{err.Error()},
			StatusCode: fiber.StatusBadRequest,
		})
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	userService := service.UserService{}
	userId, err := userService.GetUserIdByUsername(username)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			ErrorMsgs:  []string{err.Error()},
			StatusCode: fiber.StatusBadRequest,
		})
	}

	// Check whether shortened URL exists or not
	// TODO: Bind shortened url for particular user
	service := service.UrlService{}

	response, err := service.ShortenUrl(*body, userId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			ErrorMsgs:  []string{err.Error()},
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)

}

// @Summary Resolve Short URL
// @Description Resolves a short URL to its original URL and redirects the user
// @Tags Url
// @Produce json
// @Param Authorization header string true "Bearer Token in the format 'Bearer <token>'"
// @Param short path string true "Shortened URL identifier"
// @Success 301 {string} string "Redirects to the original URL"
// @Failure 400 {object} dto.ErrorResponse "Bad Request"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /{short} [get]
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
