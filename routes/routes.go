package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/WebServ/pkg"
	"github.com/svemat01/WebServ/redisDB"
	"log"
	"strconv"
	"time"
)

type countData struct {
	Device string `json:"device"`
	Value  string `json:"value"`
}

func SetupRoutes(app *fiber.App) error {
	app.Get("/", func(c *fiber.Ctx) error {

		return c.Render("index", fiber.Map{
			"Device": "test",
		})

	})

	app.Post("/api/count", func(c *fiber.Ctx) error {
		contentType := c.Get("Content-Type")

		if contentType != "application/json" {
			return pkg.BadRequest("Content-Type must be application/json")
		}

		data := new(countData)

		if err := c.BodyParser(data); err != nil {
			return pkg.BadRequest(err.Error())
		}

		log.Printf("%+v", data)

		now := time.Now().UnixMilli()
		nowString := strconv.FormatInt(now, 10)

		_, err := redisDB.RedisClient.HSet(redisDB.RedisContext, nowString, "device", data.Device, "value", data.Value).Result()

		if err != nil {
			return pkg.BadRequest(err.Error())
		}

		return c.SendStatus(fiber.StatusOK)
	})

	return nil
}
