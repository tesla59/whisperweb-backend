package confession

import (
	"time"

	"github.com/gofiber/fiber/v2"
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
	return c.SendString("Get a confession")
}

func GetAllConfession(c *fiber.Ctx) error {
	return c.SendString("Get all confessions")
}

func DeleteConfession(c *fiber.Ctx) error {
	return c.SendString("Delete a confession")
}
