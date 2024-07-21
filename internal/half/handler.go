package half

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/splorg/compstats-api/internal/config"
	"github.com/splorg/compstats-api/internal/database"
	"github.com/splorg/compstats-api/internal/util"
)

type halfHandler struct {
	*config.ApiConfig
}

func NewHalfHandler(cfg *config.ApiConfig) *halfHandler {
	return &halfHandler{ApiConfig: cfg}
}

func (h *halfHandler) CreateHalf(c *fiber.Ctx) error {
	var req createHalfDTO

	err := util.ValidateRequestBody(c, &req)
	if err != nil {
		return err
	} 

	tournament, err := h.DB.GetTournamentByID(c.Context(), int64(req.TournamentID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "tournament not found"})
	}

	half, err := h.DB.CreateHalf(c.Context(), database.CreateHalfParams{
		MapName:      req.MapName,
		Team1:        req.Team1,
		Team2:        req.Team2,
		TournamentID: tournament.ID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save half"})
	}

	return c.Status(fiber.StatusOK).JSON(half)
}

func (h *halfHandler) GetAllHalfs(c *fiber.Ctx) error {
	halfs, err := h.DB.GetAllHalfs(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get halfs"})
	}
	if halfs == nil {
		halfs = []database.Half{}
	}

	return c.Status(fiber.StatusOK).JSON(halfs)
}

func (h *halfHandler) GetHalfByID(c *fiber.Ctx) error {
	halfId, err := util.GetIDFromParams(c)
	if err != nil {
		return err
	}

	half, err := h.DB.FindHalfByID(c.Context(), int64(halfId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "half not found"})
	}

	return c.Status(fiber.StatusOK).JSON(half)
}

func (h *halfHandler) UpdateHalf(c *fiber.Ctx) error {
	halfId, err := util.GetIDFromParams(c)
	if err != nil {
		return err
	}

	var req updateHalfDTO

	err = util.ValidateRequestBody(c, &req)
	if err != nil {
		return err
	}

	half, err := h.DB.FindHalfByID(c.Context(), int64(halfId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "half not found"})
	}

	updatedHalf, err := h.DB.UpdateHalfByID(c.Context(), database.UpdateHalfByIDParams{
		MapName: sql.NullString{String: req.MapName, Valid: req.MapName != ""},
		Team1: sql.NullString{String: req.Team1, Valid: req.Team1 != ""},
		Team2: sql.NullString{String: req.Team2, Valid: req.Team2 != ""},
		ID: half.ID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update half"})
	}

	return c.Status(fiber.StatusOK).JSON(updatedHalf)
}

func (h *halfHandler) DeleteHalf(c *fiber.Ctx) error {
	halfId, err := util.GetIDFromParams(c)
	if err != nil {
		return err
	}

	half, err := h.DB.FindHalfByID(c.Context(), int64(halfId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "half not found"})
	}

	err = h.DB.DeleteHalfByID(c.Context(), half.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete half"})
	}

	return c.SendStatus(fiber.StatusOK)
}
