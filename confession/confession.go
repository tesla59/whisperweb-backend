package confession

import (
	"github.com/gofiber/fiber/v2"
)

func NewConfession(c *fiber.Ctx) error

func GetConfession(c *fiber.Ctx) error

func GetAllConfession(c *fiber.Ctx) error

func DeleteConfession(c *fiber.Ctx) error
