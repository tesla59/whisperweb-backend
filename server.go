package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tesla59/whisperweb-backend/confession"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")
	confessionRoute := v1.Group("/confession")

	confessionRoute.Post("/new", confession.NewConfession)
	confessionRoute.Get("/get/:id", confession.GetConfession)
	confessionRoute.Get("/getall", confession.GetAllConfession)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":5000")
}
