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

// @Summary Create a new user
// @Description Creates a user account with a username, email, and password
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.UserRequestDTO true "User details"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/user/register [post]
func CreateUser(c *fiber.Ctx) error {
	userDTO := new(dto.UserRequestDTO)

	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			ErrorMsgs:  []string{"could not parse the request body"},
			StatusCode: fiber.StatusBadRequest,
		})
	}

	errorMsgs := validator.RegisterUserValidator(userDTO)

	if len(errorMsgs["errors"]) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			ErrorMsgs:  errorMsgs["errors"],
			StatusCode: fiber.StatusBadRequest,
		})
	}

	service := service.UserService{}
	msg, err := service.CreateUser(userDTO)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			ErrorMsgs:  []string{"Some error occured when  creating user"},
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.SuccessResponse{
		Message:    msg,
		StatusCode: fiber.StatusCreated,
	})

}

// @Summary User Login
// @Description User logs in and a token is returned
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.UserLoginRequestDTO true "Login details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/user/login [post]
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
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
			ErrorMsgs:  []string{"unauthorized request"},
			StatusCode: fiber.StatusUnauthorized,
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
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			ErrorMsgs:  []string{err.Error()},
			StatusCode: fiber.StatusUnauthorized,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": t,
	})
}

// @Summary Get user details by username
// @Description Getting user details from username
// @Tags User
// @Produce json
// @Param Authorization header string true "Bearer Token in the format 'Bearer <token>'"
// @Param username path string true "Username of the user"
// @Success 200 {object} dto.UserResponseDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/user/{username} [get]
func GetUserByUserName(c *fiber.Ctx) error {
	username := c.Params("username")

	service := service.UserService{}

	response, err := service.GetUserByUserName(username)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			ErrorMsgs:  []string{err.Error()},
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary Delete user by username
// @Description Delete user from username
// @Tags User
// @Produce json
// @Param Authorization header string true "Bearer Token in the format 'Bearer <token>'"
// @Param username path string true "Username of the user"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} dto.ErrorResponse
// @Router /api/v1/user/{username} [delete]
func DeleteUserByUserName(c *fiber.Ctx) error {
	username := c.Params("username")

	service := service.UserService{}

	msg, err := service.DeleteUserByUserName(username)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			ErrorMsgs:  []string{err.Error()},
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": msg,
	})
}
