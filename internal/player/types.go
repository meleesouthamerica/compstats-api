package player

type createDTO struct {
	Name      string `json:"name" validate:"required"`
	VirtualID string `json:"virtual_id" validate:"required"`
}

type updateDTO struct {
	Name      string `json:"name"`
	VirtualID string `json:"virtual_id"`
}

type playerStats struct {
	Kills   int  `json:"kills" validate:"required"`
	Deaths  int  `json:"deaths" validate:"required"`
	Assists int  `json:"assists" validate:"required"`
	Score   int  `json:"score" validate:"required"`
	Winner  bool `json:"win" validate:"required"`
}

type playerToBeUpdated struct {
	Name      string      `json:"name" validate:"required"`
	VirtualID string      `json:"virtual_id" validate:"required"`
	Stats     playerStats `json:"stats" validate:"required"`
}

type updateStatsDTO struct {
	MapName        string              `json:"mapName" validate:"required"`
	Team1          string              `json:"team1" validate:"required"`
	Team2          string              `json:"team2" validate:"required"`
	GameType       string              `json:"gameType" validate:"required"`
	TournamentName string              `json:"tournamentName" validate:"required"`
	Players        []playerToBeUpdated `json:"players" validate:"required"`
}
