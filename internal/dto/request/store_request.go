package request

type CreateStoreRequestDTO struct {
	Name          string `json:"name" binding:"required"`
	Address       string `json:"address" binding:"required"`
	ParentStoreID string `json:"parent_store_id" binding:"required,uuid"`
}

type UpdateStoreDTO struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}
