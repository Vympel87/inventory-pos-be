package request

import (
	"github.com/google/uuid"
)

type CreateStoreRequestDTO struct {
	Name          string    `json:"name" binding:"required"`
	Address       string    `json:"address" binding:"required"`
	ParentStoreID uuid.UUID `json:"parent_store_id" binding:"required"`
}

type UpdateStoreDTO struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}
