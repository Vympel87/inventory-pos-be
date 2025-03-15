package domain

import (
	"time"

	enum "github.com/Vympel87/inventory-pos-be/internal/constants"
	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID              `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	StoreID       uuid.UUID              `gorm:"type:uuid"`
	Store         Store                  `gorm:"foreignKey:StoreID"`
	Name          string                 `gorm:"type:varchar(255);not null"`
	Email         string                 `gorm:"type:varchar(255);unique;not null"`
	Password      string                 `gorm:"type:text;not null"`
	Role          enum.RoleTypeEnum      `gorm:"type:role_enum;not null"`
	AccountStatus enum.AccountStatusEnum `gorm:"type:account_status_enum;not null;default:'ACTIVE'"`
	LastLogin     *time.Time
	Sessions      []Session      `gorm:"foreignKey:UserID"`
	ProductOrders []ProductOrder `gorm:"foreignKey:UserID"`
	Sales         []Sale         `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
}
