package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/shreekumar2901/url-shortener/routes"
)

func hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Hello",
	})
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/hello", hello)
	app.Post("/api/v1/shorten/me", routes.ShortenUrl)
	app.Get("/:url", routes.ResolveUrl)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
