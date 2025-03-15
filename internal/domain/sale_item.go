package domain

import (
	"time"
)

type SaleItem struct {
	ID         int     `gorm:"primaryKey;autoIncrement"`
	SaleID     int     `gorm:"not null"`
	ProductID  int     `gorm:"not null"`
	Sale       Sale    `gorm:"foreignKey:SaleID"`
	Product    Product `gorm:"foreignKey:ProductID"`
	Quantity   int     `gorm:"not null"`
	UnitPrice  float64 `gorm:"type:decimal(10,2);not null"`
	TotalPrice float64 `gorm:"type:decimal(10,2);not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
