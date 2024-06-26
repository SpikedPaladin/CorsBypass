package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 1024,
	})

	app.Use(logger.New())

	app.Use(cors.New())

	app.Get("/bypass", func(c *fiber.Ctx) error {
		url := c.Query("url")

		resp, _ := http.Get(url)

		return c.SendStream(resp.Body)
	})

	log.Fatal(app.Listen(":5000"))
}
