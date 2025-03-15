package response

import (
	"time"

	"github.com/google/uuid"
)

type UserCompleteResponseDTO struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Role          string    `json:"role"`
	AccountStatus string    `json:"account_status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type StoreSimpleDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	StoreType string    `json:"store_type"`
}

type UserWithStoreResponseDTO struct {
	ID            uuid.UUID       `json:"id"`
	Name          string          `json:"name"`
	Email         string          `json:"email"`
	Role          string          `json:"role"`
	AccountStatus string          `json:"account_status"`
	Store         *StoreSimpleDTO `json:"store"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}
