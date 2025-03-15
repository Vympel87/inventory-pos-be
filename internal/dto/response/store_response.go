package response

import (
	"time"

	"github.com/google/uuid"
)

type StoreResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	StoreType string    `json:"store_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CompleteStoreResponseDTO struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Address       string    `json:"address"`
	Phone         string    `json:"phone"`
	StoreType     string    `json:"store_type"`
	ParentStoreID uuid.UUID `json:"parent_store_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
