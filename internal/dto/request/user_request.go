package request

import (
	enum "github.com/Vympel87/inventory-pos-be/internal/constants"
	"github.com/google/uuid"
)

type CreateUserRequestDTO struct {
	StoreID  uuid.UUID `json:"store_id" validate:"required"`
	Name     string    `json:"name" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=6"`
	Role     string    `json:"role" validate:"required,oneof=SUPERVISOR ADMIN"`
}

type UpdateUserRequestDTO struct {
	Name          string                 `json:"name" validate:"required"`
	Email         string                 `json:"email" validate:"required,email"`
	Role          enum.RoleTypeEnum      `json:"role" validate:"required,oneof=ADMIN CASHIER STAFF"`
	AccountStatus enum.AccountStatusEnum `json:"account_status" validate:"required,oneof=ACTIVE INACTIVE SUSPENDED"`
	StoreID       uuid.UUID              `json:"store_id" validate:"required"`
}
