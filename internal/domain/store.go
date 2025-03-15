package domain

import (
	"time"

	enum "github.com/Vympel87/inventory-pos-be/internal/constants"
	"github.com/google/uuid"
)

type Store struct {
	ID            uuid.UUID          `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name          string             `gorm:"type:varchar(255);not null"`
	Address       string             `gorm:"type:text"`
	Phone         string             `gorm:"type:varchar(50)"`
	Email         string             `gorm:"type:varchar(255)"`
	StoreType     enum.StoreTypeEnum `gorm:"type:store_type_enum;not null"`
	ParentStoreID *uuid.UUID         `gorm:"type:uuid"`
	ParentStore   *Store             `gorm:"foreignKey:ParentStoreID"`
	ChildStores   []Store            `gorm:"foreignKey:ParentStoreID"`
	Users         []User             `gorm:"foreignKey:StoreID"`
	Inventories   []Inventory        `gorm:"foreignKey:StoreID"`
	ProductOrders []ProductOrder     `gorm:"foreignKey:StoreID"`
	Sales         []Sale             `gorm:"foreignKey:StoreID"`
	CreatedAt     time.Time          `gorm:"autoCreateTime"`
	UpdatedAt     time.Time          `gorm:"autoUpdateTime"`
}
