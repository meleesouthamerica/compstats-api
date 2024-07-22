package tournament

import (
	"github.com/gofiber/fiber/v2"
	"github.com/splorg/compstats-api/internal/config"
	"github.com/splorg/compstats-api/internal/database"
	"github.com/splorg/compstats-api/internal/util"
	"github.com/splorg/compstats-api/internal/validator"
)

type tournamentHandler struct {
	*config.ApiConfig
}

func NewTournamentHandler(cfg *config.ApiConfig) *tournamentHandler {
	return &tournamentHandler{ApiConfig: cfg}
}

// GetAllTournaments godoc
//
//	@Summary		Tournaments
//	@Description	get all tournaments
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{array}	database.Tournament
//	@Router			/admin/tournaments [get]
func (h *tournamentHandler) GetAllTournaments(c *fiber.Ctx) error {
	tournaments, err := h.DB.GetAllTournaments(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if tournaments == nil {
		tournaments = []database.Tournament{}
	}

	return c.Status(fiber.StatusOK).JSON(tournaments)
}

// GetTournamentByID godoc
//
//	@Summary		Get a tournament
//	@Description	get tournament by ID
//	@Produce		json
//	@Param			id	path		int	true	"Tournament ID"
//	@Success		200	{object}	database.Tournament
//	@Router			/admin/tournaments/{id} [get]
func (h *tournamentHandler) GetTournamentByID(c *fiber.Ctx) error {
	tournamentId, err := util.GetIDFromParams(c)
	if err != nil {
		return err
	}

	tournament, err := h.DB.GetTournamentByID(c.Context(), int64(tournamentId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "tournament not found"})
	}

	return c.Status(fiber.StatusOK).JSON(tournament)
}

// CreateTournament godoc
//
//	@Summary		Create tournament
//	@Description	create a tournament
//	@Accept			json
//	@Produce		json
//
// @Param dto body tournament.createUpdateDTO true "create json"
//
//	@Success		201	{object}	database.Tournament
//	@Router			/admin/tournaments [post]
func (h *tournamentHandler) CreateTournament(c *fiber.Ctx) error {
	var req createUpdateDTO

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newTournament, err := h.DB.CreateTournament(c.Context(), req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "tournament name must be unique"})
	}

	return c.Status(fiber.StatusCreated).JSON(newTournament)
}

// UpdateTournament godoc
//
//	@Summary		Update a tournament
//	@Description	update tournament by id
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Tournament ID"
//	@Param			dto	body		tournament.createUpdateDTO	true	"update json"
//	@Success		200		{object}	database.Tournament
//	@Router			/admin/tournaments/{id} [patch]
func (h *tournamentHandler) UpdateTournament(c *fiber.Ctx) error {
	tournamentId, err := util.GetIDFromParams(c)
	if err != nil {
		return err
	}

	var req createUpdateDTO

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	tournament, err := h.DB.GetTournamentByID(c.Context(), int64(tournamentId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "tournament not found"})
	}

	updatedTournament, err := h.DB.UpdateTournamentByID(c.Context(), database.UpdateTournamentByIDParams{
		Name: req.Name,
		ID:   tournament.ID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(updatedTournament)
}

// DeleteTournament godoc
//
//	@Summary		Delete a tournament
//	@Description	delete tournament by id
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Tournament ID"
//	@Router			/admin/tournaments/{id} [delete]
func (h *tournamentHandler) DeleteTournament(c *fiber.Ctx) error {
	tournamentId, err := util.GetIDFromParams(c)
	if err != nil {
		return err
	}

	tournament, err := h.DB.GetTournamentByID(c.Context(), int64(tournamentId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "tournament not found"})
	}

	err = h.DB.DeleteTournamentByID(c.Context(), tournament.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete tournament"})
	}

	return c.SendStatus(fiber.StatusOK)
}
