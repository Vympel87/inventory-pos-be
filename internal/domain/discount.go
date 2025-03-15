package domain

import (
	"time"

	enum "github.com/Vympel87/inventory-pos-be/internal/constants"
)

type Discount struct {
	ID            int                   `gorm:"primaryKey;autoIncrement"`
	Name          string                `gorm:"size:255;not null"`
	DiscountType  enum.DiscountTypeEnum `gorm:"type:discount_type_enum;not null"`
	DiscountValue float64               `gorm:"type:decimal(10,2);not null"`
	StartDate     *time.Time
	EndDate       *time.Time
	SaleDiscounts []SaleDiscount
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
