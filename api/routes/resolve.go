package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/shreekumar2901/url-shortener/database"
)

func ResolveUrl(c *fiber.Ctx) error {
	url := c.Params("url")

	rdb := database.CreateDBClient(0)
	defer rdb.Close()

	value, err := rdb.Get(database.Ctx, url).Result()

	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found in the database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	}

	rIncr := database.CreateDBClient(1)
	defer rIncr.Close()

	_ = rIncr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)

}
