package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/shreekumar2901/url-shortener/routes"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/v1/shorten", routes.ShortenUrl)
	app.Get("/:url", routes.ResolveUrl)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	setupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "API does not exist",
		})
	})

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
