package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/tesla59/whisperweb-backend/confession"
	"github.com/tesla59/whisperweb-backend/database"
)

func initDB() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("confessions.db"))
	if err != nil {
		panic("Cannot connect to Database")
	}
	fmt.Println("Connected to DB")
	database.DBConn.AutoMigrate(&confession.Confession{})
	fmt.Println("Migrated successfully")
}

func main() {
	app := fiber.New()
	initDB()

	api := app.Group("/api")
	v1 := api.Group("/v1")
	confessionRoute := v1.Group("/confession")

	confessionRoute.Post("/new", confession.NewConfession)
	confessionRoute.Get("/get/:id", confession.GetConfession)
	confessionRoute.Get("/getall", confession.GetAllConfession)
	confessionRoute.Delete("/delete/:id", confession.DeleteConfession)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":5000")
}
