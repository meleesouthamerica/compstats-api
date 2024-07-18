package tournament

type createUpdateDTO struct {
	Name     string `json:"name" validate:"required"`
}