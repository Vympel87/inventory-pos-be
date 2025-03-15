package domain

import (
	"time"
)

type SaleDiscount struct {
	ID         int      `gorm:"primaryKey;autoIncrement"`
	SaleID     int      `gorm:"not null"`
	DiscountID int      `gorm:"not null"`
	Sale       Sale     `gorm:"foreignKey:SaleID"`
	Discount   Discount `gorm:"foreignKey:DiscountID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
