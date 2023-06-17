package confession

import (
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
	return c.SendString("New Confession")
}

func GetConfession(c *fiber.Ctx) error {
	db := database.DBConn
	var confession Confession
	db.First(&confession)
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
