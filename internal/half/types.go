package half

type createDTO struct {
	MapName      string `json:"mapName" validate:"required"`
	Team1        string `json:"team1" validate:"required"`
	Team2        string `json:"team2" validate:"required"`
	GameType     string `json:"gameType" validate:"required"`
	TournamentID int    `json:"tournamentId" validate:"required"`
}

type updateDTO struct {
	MapName      string `json:"mapName"`
	Team1        string `json:"team1"`
	Team2        string `json:"team2"`
}
