package confession

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tesla59/whisperweb-backend/database"
)

type Confession struct {
	ID        string    `json:"id"`
	By        string    `json:"by"`
	To        string    `json:"to"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewConfession(c *fiber.Ctx) error {
	db := database.DBConn
	var confession Confession
	if err := c.BodyParser(&confession); err != nil {
		c.Status(http.StatusBadRequest).SendString(fmt.Sprint(err))
	}
	db.Create(&confession)
	return c.Status(http.StatusOK).JSON(confession)
}

func GetConfession(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var confession Confession
	db.First(&confession, id)
	if confession.ID == "" {
		return c.Status(http.StatusBadRequest).JSON("not found")
	}
	return c.JSON(confession)
}

func GetAllConfession(c *fiber.Ctx) error {
	db := database.DBConn
	var confessions []Confession
	db.Find(&confessions)
	return c.JSON(confessions)
}

func DeleteConfession(c *fiber.Ctx) error {
	return c.SendString("Delete a confession")
}
