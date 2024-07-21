package player

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/splorg/compstats-api/internal/config"
	"github.com/splorg/compstats-api/internal/database"
	"github.com/splorg/compstats-api/internal/util"
	"github.com/splorg/compstats-api/internal/validator"
)

type playerHandler struct {
	*config.ApiConfig
}

func NewPlayerHandler(cfg *config.ApiConfig) *playerHandler {
	return &playerHandler{ApiConfig: cfg}
}

func (h *playerHandler) CreatePlayer(c *fiber.Ctx) error {
	var req createDTO

	err := util.ValidateRequestBody(c, &req)
	if err != nil {
		return err
	}

	player, err := h.DB.CreatePlayer(c.Context(), database.CreatePlayerParams{
		Name:      req.Name,
		VirtualID: req.VirtualID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save player"})
	}

	return c.Status(fiber.StatusCreated).JSON(player)
}

func (h *playerHandler) GetAllPlayers(c *fiber.Ctx) error {
	players, err := h.DB.GetAllPlayers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get players"})
	}
	if players == nil {
		players = []database.Player{}
	}

	return c.Status(fiber.StatusOK).JSON(players)
}

func (h *playerHandler) GetPlayerByID(c *fiber.Ctx) error {
	playerId, err := util.GetIDFromParams(c)
	if err != nil {
		return err
	}

	player, err := h.DB.FindPlayerByID(c.Context(), int64(playerId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "player not found"})
	}

	return c.Status(fiber.StatusOK).JSON(player)
}

func (h *playerHandler) UpdatePlayer(c *fiber.Ctx) error {
	playerId, err := util.GetIDFromParams(c)
	if err != nil {
		return err
	}

	var req updateDTO

	err = util.ValidateRequestBody(c, &req)
	if err != nil {
		return err
	}

	player, err := h.DB.FindPlayerByID(c.Context(), int64(playerId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "player not found"})
	}

	updatedPlayer, err := h.DB.UpdatePlayerByID(c.Context(), database.UpdatePlayerByIDParams{
		Name:      sql.NullString{String: req.Name, Valid: req.Name != ""},
		VirtualID: sql.NullString{String: req.VirtualID, Valid: req.VirtualID != ""},
		ID:        player.ID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(updatedPlayer)
}

func (h *playerHandler) DeletePlayer(c *fiber.Ctx) error {
	playerId, err := util.GetIDFromParams(c)
	if err != nil {
		return err
	}

	player, err := h.DB.FindPlayerByID(c.Context(), int64(playerId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "player not found"})
	}

	err = h.DB.DeletePlayerByID(c.Context(), player.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete player"})
	}

	return c.SendStatus(fiber.StatusOK)
}

// save half and batch update player stats after game mission ends
func (h *playerHandler) UpdatePlayerStats(c *fiber.Ctx) error {
	var req updateStatsDTO

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	currentTournament, err := h.DB.GetTournamentByName(c.Context(), req.TournamentName)
	if err != nil {
		currentTournament, err = h.DB.CreateTournament(c.Context(), req.TournamentName)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "failed to save tournament"})
		}
	}

	half, err := h.DB.CreateHalf(c.Context(), database.CreateHalfParams{
		MapName:      req.MapName,
		Team1:        req.Team1,
		Team2:        req.Team2,
		TournamentID: currentTournament.ID,
	})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "failed to save half"})
	}

	for _, player := range req.Players {
		existingPlayer, err := h.DB.FindPlayerByVirtualID(c.Context(), player.VirtualID)
		if err != nil {
			existingPlayer, err = h.DB.CreatePlayer(c.Context(), database.CreatePlayerParams{
				Name:      player.Name,
				VirtualID: player.VirtualID,
			})
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
		}

		_, err = h.DB.CreateStats(c.Context(), database.CreateStatsParams{
			PlayerID: existingPlayer.ID,
			HalfID:   half.ID,
			Kills:    int64(player.Stats.Kills),
			Deaths:   int64(player.Stats.Deaths),
			Assists:  int64(player.Stats.Assists),
			Score:    int64(player.Stats.Score),
			Winner:   player.Stats.Winner,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "players stats updated successfully"})
}
