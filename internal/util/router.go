package util

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/splorg/compstats-api/internal/validator"
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

func ValidateRequestBody(c *fiber.Ctx, req interface{}) error {
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return nil
}
