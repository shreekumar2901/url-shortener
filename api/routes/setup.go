package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, jwt *fiber.Handler) {
	// User
	app.Post("/api/v1/user/register", CreateUser)
	app.Post("/api/v1/user/login", Login)
	app.Get("/api/v1/user/:username", *jwt, GetUserByUserName)
	app.Delete("/api/v1/user/:username", *jwt, DeleteUserByUserName)

	// Url
	app.Post("/api/v1/shorten", *jwt, ShortenUrl)
	app.Get("/api/v1/urls", *jwt, ListUrls)
	app.Delete("/api/v1/urls", *jwt, DeleteShortByUrl)
	app.Get("/:short", *jwt, ResolveUrl)
}
