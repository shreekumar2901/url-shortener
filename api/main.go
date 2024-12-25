package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/shreekumar2901/url-shortener/config"
	"github.com/shreekumar2901/url-shortener/database"
	_ "github.com/shreekumar2901/url-shortener/docs"
	"github.com/shreekumar2901/url-shortener/middlewares"
	"github.com/shreekumar2901/url-shortener/routes"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title URL Shortener API
// @version 1.0
// @description This is a REST API for a URL shortener service.
// @host localhost:3000
func main() {
	database.Connect()
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	jwt := middlewares.NewAuthMiddleware(config.Config("JWT_SECRET"))

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	routes.SetupRoutes(app, &jwt)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "API does not exist",
		})
	})

	app_port := config.Config("APP_PORT")
	log.Fatal(app.Listen(app_port))
}
