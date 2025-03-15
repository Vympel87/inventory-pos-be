package domain

import (
	"time"

	"github.com/google/uuid"
)

type Sale struct {
	ID            int       `gorm:"primaryKey;autoIncrement"`
	StoreID       uuid.UUID `gorm:"type:uuid;not null"`
	UserID        *uuid.UUID
	Store         Store     `gorm:"foreignKey:StoreID"`
	User          *User     `gorm:"foreignKey:UserID"`
	SaleDate      time.Time `gorm:"not null"`
	SaleNumber    string    `gorm:"size:50;unique;not null"`
	PaymentMethod string    `gorm:"size:50"`
	TotalAmount   float64   `gorm:"type:decimal(10,2);not null"`
	SaleItems     []SaleItem
	SaleDiscounts []SaleDiscount
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
