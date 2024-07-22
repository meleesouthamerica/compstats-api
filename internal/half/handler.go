package half

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/splorg/compstats-api/internal/config"
	"github.com/splorg/compstats-api/internal/database"
	"github.com/splorg/compstats-api/internal/util"
	"github.com/splorg/compstats-api/internal/validator"
)

type halfHandler struct {
	*config.ApiConfig
}

func NewHalfHandler(cfg *config.ApiConfig) *halfHandler {
	return &halfHandler{ApiConfig: cfg}
}

// CreateHalf godoc
//
//	@Summary		Create half
//	@Description	create a half
//	@Accept			json
//	@Produce		json
//
// @Param dto body half.createDTO true "create json"
//
//	@Success		201	{object}	database.Half
//	@Router			/admin/halfs [post]
func (h *halfHandler) CreateHalf(c *fiber.Ctx) error {
	var req createDTO

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
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

// GetAllHalfs godoc
//
//	@Summary		Halfs
//	@Description	get all halfs
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{array}	database.Half
//	@Router			/admin/halfs [get]
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

// GetHalfByID godoc
//
//	@Summary		Get a half
//	@Description	get half by ID
//	@Produce		json
//	@Param			id	path		int	true	"Half ID"
//	@Success		200	{object}	database.Half
//	@Router			/admin/halfs/{id} [get]
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

// UpdateHalf godoc
//
//	@Summary		Update a half
//	@Description	update half by id
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Half ID"
//	@Param			dto	body		half.updateDTO	true	"update json"
//	@Success		200		{object}	database.Half
//	@Router			/admin/halfs/{id} [patch]
func (h *halfHandler) UpdateHalf(c *fiber.Ctx) error {
	halfId, err := util.GetIDFromParams(c)
	if err != nil {
		return err
	}

	var req updateDTO

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	half, err := h.DB.FindHalfByID(c.Context(), int64(halfId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "half not found"})
	}

	updatedHalf, err := h.DB.UpdateHalfByID(c.Context(), database.UpdateHalfByIDParams{
		MapName: sql.NullString{String: req.MapName, Valid: req.MapName != ""},
		Team1:   sql.NullString{String: req.Team1, Valid: req.Team1 != ""},
		Team2:   sql.NullString{String: req.Team2, Valid: req.Team2 != ""},
		ID:      half.ID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update half"})
	}

	return c.Status(fiber.StatusOK).JSON(updatedHalf)
}

// DeleteHalf godoc
//
//	@Summary		Delete a half
//	@Description	delete half by id
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Half ID"
//	@Router			/admin/halfs/{id} [delete]
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
