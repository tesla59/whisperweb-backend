package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/tesla59/whisperweb-backend/handlers"
	"github.com/tesla59/whisperweb-backend/database"
	"github.com/tesla59/whisperweb-backend/models"
)

func initDB() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("confessions.db"))
	if err != nil {
		panic("Cannot connect to Database")
	}
	fmt.Println("Connected to DB")
	database.DBConn.AutoMigrate(&models.Confession{})
	fmt.Println("Migrated successfully")
}

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
	initDB()
	initRoutes(app)
	app.Listen(":5000")
}
