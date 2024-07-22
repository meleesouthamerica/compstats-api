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

// CreatePlayer godoc
//
//	@Summary		Create player
//	@Description	create a player
//	@Accept			json
//	@Produce		json
//
// @Param dto body player.createDTO true "create json"
//
//	@Success		201	{object}	database.Player
//	@Router			/admin/players [post]
func (h *playerHandler) CreatePlayer(c *fiber.Ctx) error {
	var req createDTO

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
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

// GetAllPlayers godoc
//
//	@Summary		Players
//	@Description	get all players
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{array}	database.Player
//	@Router			/admin/players [get]
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

// GetPlayerByID godoc
//
//	@Summary		Get a player
//	@Description	get player by ID
//	@Produce		json
//	@Param			id	path		int	true	"Player ID"
//	@Success		200	{object}	database.Player
//	@Router			/admin/players/{id} [get]
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

// UpdatePlayer godoc
//
//	@Summary		Update a player
//	@Description	update player by id
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Player ID"
//	@Param			dto	body		player.updateDTO	true	"update json"
//	@Success		200		{object}	database.Player
//	@Router			/admin/players/{id} [patch]
func (h *playerHandler) UpdatePlayer(c *fiber.Ctx) error {
	playerId, err := util.GetIDFromParams(c)
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

// DeletePlayer godoc
//
//	@Summary		Delete a player
//	@Description	delete player by id
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Player ID"
//	@Router			/admin/players/{id} [delete]
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

// UpdatePlayerStats godoc
//
//	@Summary		Update stats
//	@Description	save half, players and stats from game server
//	@Accept			json
//	@Produce		json
//
// @Param dto body player.updateStatsDTO true "update stats json"
//
//	@Success		201	{object}	player.updateStatsResponse
//	@Router			/stats [post]
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

	return c.Status(fiber.StatusCreated).JSON(updateStatsResponse{Message: "players stats updated successfully"})
}
