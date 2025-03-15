package response

import (
	"time"

	"github.com/google/uuid"
)

type StoreAuthResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	StoreType string    `json:"store_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserAuthResponseDTO struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Role          string    `json:"role"`
	AccountStatus string    `json:"account_status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type RegisterResponseDTO struct {
	Store StoreAuthResponseDTO `json:"store"`
	User  UserAuthResponseDTO  `json:"user"`
}

type LoginResponseDTO struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	Role          string     `json:"role"`
	AccountStatus string     `json:"account_status"`
	LastLogin     *time.Time `json:"last_login,omitempty"`
}

type LogoutResponseDTO struct {
	Message string `json:"message"`
}
