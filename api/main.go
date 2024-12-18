package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/shreekumar2901/url-shortener/config"
	"github.com/shreekumar2901/url-shortener/database"
	"github.com/shreekumar2901/url-shortener/routes"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	routes.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "API does not exist",
		})
	})

	app_port := config.Config("APP_PORT")
	log.Fatal(app.Listen(app_port))
}
