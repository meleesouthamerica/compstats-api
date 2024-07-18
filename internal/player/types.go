package player

type createDTO struct {
	Name   string `json:"name" validate:"required"`
	VirtualID string `json:"virtual_id" validate:"required"`
}

type updateDTO struct {
  Name   string `json:"name" validate:"required"`
}

type playerStats struct {
  Kills int `json:"kills" validate:"required"`
  Deaths int `json:"deaths" validate:"required"`
  Assists int `json:"assists" validate:"required"`
  Score int `json:"score" validate:"required"`
  Winner bool `json:"win" validate:"required"`
}

type playerToBeUpdated struct {
  Name   string `json:"name" validate:"required"`
  VirtualID string `json:"virtual_id" validate:"required"`
  Stats []playerStats `json:"stats" validate:"required"`
}

type updateStatsDTO struct {
  Players []playerToBeUpdated `json:"players" validate:"required"`
  HalfID int `json:"matchId" validate:"required"`
}
