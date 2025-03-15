package domain

import (
	"time"
)

type Product struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	CategoryID  *int
	Category    *Category `gorm:"foreignKey:CategoryID"`
	Name        string    `gorm:"size:255;not null"`
	Description string
	SKU         string `gorm:"size:50;unique"`
	ImageURL    string
	Size        string  `gorm:"size:5"`
	Color       string  `gorm:"size:30"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	Inventories []Inventory
	POItems     []POItem
	SaleItems   []SaleItem
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
