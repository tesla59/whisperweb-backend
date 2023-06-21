package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tesla59/whisperweb-backend/database"
	"github.com/tesla59/whisperweb-backend/models"
)

func NewConfession(c *fiber.Ctx) error {
	db := database.DBConn
	var confession models.Confession
	if err := c.BodyParser(&confession); err != nil {
		c.Status(http.StatusBadRequest).SendString(fmt.Sprint(err))
		return err
	}
	confession.ID = uuid.NewString()
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
