package confession

import (
	"github.com/gofiber/fiber/v2"
)

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
