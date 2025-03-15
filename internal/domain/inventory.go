package domain

import (
	"time"

	"github.com/google/uuid"
)

type Inventory struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	ProductID     uint      `gorm:"type:integer"`
	Product       Product   `gorm:"foreignKey:ProductID"`
	StoreID       uuid.UUID `gorm:"type:uuid"`
	Store         Store     `gorm:"foreignKey:StoreID"`
	Quantity      int       `gorm:"not null;default:0"`
	MinStockLevel int       `gorm:"not null;default:0"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
