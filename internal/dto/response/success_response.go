package response

type SuccessResponse[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type SuccessPaginatedResponse[T any, P any] struct {
	Status     string `json:"status"`
	Data       []T    `json:"data"`
	Pagination P      `json:"pagination"`
}

type PaginationMeta struct {
	TotalItems   int `json:"total_items"`
	TotalPages   int `json:"total_pages"`
	CurrentPage  int `json:"current_page"`
	ItemsPerPage int `json:"items_per_page"`
}

type SuccessDeleteResponseDTO struct {
	Message string `json:"message"`
}
