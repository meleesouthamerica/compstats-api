package tournament

import (
	"github.com/gofiber/fiber/v2"
	"github.com/splorg/compstats-api/internal/config"
	"github.com/splorg/compstats-api/internal/database"
	"github.com/splorg/compstats-api/internal/util"
)

type tournamentHandler struct {
	*config.ApiConfig
}

func NewTournamentHandler(cfg *config.ApiConfig) *tournamentHandler {
	return &tournamentHandler{ApiConfig: cfg}
}

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

func (h *tournamentHandler) CreateTournament(c *fiber.Ctx) error {
	var req createUpdateDTO

	err := util.ValidateRequestBody(c, &req)
	if err != nil {
		return err
	} 

	newTournament, err := h.DB.CreateTournament(c.Context(), req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "tournament name must be unique"})
	}

	return c.Status(fiber.StatusCreated).JSON(newTournament)
}

func (h *tournamentHandler) UpdateTournament(c *fiber.Ctx) error {
	tournamentId, err := util.GetIDFromParams(c)
	if err != nil {
		return err
	}

	var req createUpdateDTO

	err = util.ValidateRequestBody(c, &req)
	if err != nil {
		return err
	} 

	tournament, err := h.DB.GetTournamentByID(c.Context(), int64(tournamentId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "tournament not found"})
	}

	updatedTournament, err := h.DB.UpdateTournamentByID(c.Context(), database.UpdateTournamentByIDParams{
		Name:      req.Name,
		ID:        tournament.ID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(updatedTournament)
}

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
