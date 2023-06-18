package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tesla59/whisperweb-backend/database"
	"github.com/tesla59/whisperweb-backend/handlers"
)

func initRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	confessionRoute := v1.Group("/confession")

	confessionRoute.Post("/new", handlers.NewConfession)
	confessionRoute.Get("/get/:id", handlers.GetConfession)
	confessionRoute.Get("/getall", handlers.GetAllConfession)
	confessionRoute.Delete("/delete/:id", handlers.DeleteConfession)
}

func main() {
	app := fiber.New()
	database.InitDB()
	initRoutes(app)
	app.Listen(":5000")
}
