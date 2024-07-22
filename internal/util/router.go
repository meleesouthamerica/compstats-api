package util

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetIDFromParams(c *fiber.Ctx) (int, error) {
	idParam := c.Params("id")
	if idParam == "" {
		return 0, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID must be a number"})
	}

	return id, nil
}