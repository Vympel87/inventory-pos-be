package request

type CategoryCompleteRequestDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
