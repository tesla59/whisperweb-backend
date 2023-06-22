package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tesla59/whisperweb-backend/database"
	"github.com/tesla59/whisperweb-backend/models"
)

func InitRoutes(app *fiber.App) {
	app.Use(cors.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")
	confessionRoute := v1.Group("/confession")

	confessionRoute.Post("/new", NewConfession)
	confessionRoute.Get("/get/:id", GetConfession)
	confessionRoute.Get("/getall", GetAllConfession)
	confessionRoute.Delete("/delete/:id", DeleteConfession)
}

func NewConfession(c *fiber.Ctx) error {
	db := database.DBConn
	var confession models.Confession
	if err := c.BodyParser(&confession); err != nil {
		c.Status(http.StatusBadRequest).SendString(fmt.Sprint(err))
		return err
	}
	confession.ID = uuid.NewString()
	confession.CreatedAt = time.Now()
	confession.UpdatedAt = time.Now()
	// Check unique key constraint
	db.Create(&confession)
	return c.Status(http.StatusOK).JSON(confession)
}

func GetConfession(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var confession models.Confession
	db.First(&confession, id)
	if confession.ID == "" {
		return c.Status(http.StatusBadRequest).JSON("not found")
	}
	return c.JSON(confession)
}

func GetAllConfession(c *fiber.Ctx) error {
	db := database.DBConn
	var confessions []models.Confession
	db.Find(&confessions)
	return c.JSON(confessions)
}

func DeleteConfession(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var confession models.Confession
	db.First(&confession, id)
	if confession.ID == "" {
		return c.Status(http.StatusBadRequest).JSON("not found")
	}
	db.Delete(&confession)
	return c.JSON(confession)
}
