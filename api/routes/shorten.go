package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/shreekumar2901/url-shortener/database"
	"github.com/shreekumar2901/url-shortener/helpers"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"customShort"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"customShort"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenUrl(c *fiber.Ctx) error {

	// acts like an object creation
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Can not parse body",
		})
	}

	rdb := database.CreateDBClient(1)
	defer rdb.Close()

	val, err := rdb.Get(database.Ctx, c.IP()).Result()

	if err == redis.Nil {
		_ = rdb.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		val, _ = rdb.Get(database.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(val)

		if valInt <= 0 {
			limit, _ := rdb.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"err":        "Rate limit exceeded",
				"rate_limit": limit,
			})

		}
	}

	// Check for valid URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	if helpers.DetectDomainError(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "This domain is not possible :)",
		})
	}

	body.URL = helpers.EnforeHTTP(body.URL)

	var id string

	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	r := database.CreateDBClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, id).Result()

	if val != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Can not use given custom short",
		})
	}

	if body.Expiry == 0 {
		body.Expiry = 48
	}

	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Some error occurred when connecting to database",
		})
	}

	resp := response{
		URL:             body.URL,
		CustomShort:     "",
		Expiry:          body.Expiry,
		XRateRemaining:  10,
		XRateLimitReset: 30,
	}

	rdb.Decr(database.Ctx, c.IP())

	val, _ = rdb.Get(database.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := rdb.TTL(database.Ctx, c.IP()).Result()
	resp.XRateLimitReset = ttl

	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id

	return c.Status(fiber.StatusOK).JSON(resp)
}
