package domain

import (
	"time"

	enum "github.com/Vympel87/inventory-pos-be/internal/constants"
	"github.com/google/uuid"
)

type ProductOrder struct {
	ID          int                  `gorm:"primaryKey;autoIncrement"`
	StoreID     uuid.UUID            `gorm:"type:uuid;not null"`
	UserID      uuid.UUID            `gorm:"type:uuid;not null"`
	Store       Store                `gorm:"foreignKey:StoreID"`
	User        User                 `gorm:"foreignKey:UserID"`
	OrderDate   time.Time            `gorm:"not null"`
	OrderNumber string               `gorm:"size:50;unique;not null"`
	Status      enum.OrderStatusEnum `gorm:"type:order_status_enum;default:'PENDING'"`
	Notes       string
	POItems     []POItem `gorm:"foreignKey:POID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
