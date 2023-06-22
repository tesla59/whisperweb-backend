package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tesla59/whisperweb-backend/database"
	"github.com/tesla59/whisperweb-backend/handlers"
)

func main() {
	app := fiber.New()
	database.InitDB()
	handlers.InitRoutes(app)
	app.Listen(":5000")
}
