package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Post("/api/v1/user/register", CreateUser)
	app.Get("/api/v1/user/:username", GetUserByUserName)
	app.Delete("/api/v1/user/:username", DeleteUserByUserName)

	app.Post("/api/v1/shorten", ShortenUrl)
	app.Get("/:url", ResolveUrl)
}
