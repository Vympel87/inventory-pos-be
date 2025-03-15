package response

import (
	"time"

	"github.com/google/uuid"
)

type InventoryWithProductDTO struct {
	ID            uint              `json:"id"`
	StoreID       uuid.UUID         `json:"store_id"`
	Quantity      int               `json:"quantity"`
	MinStockLevel int               `json:"min_stock_level"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	Product       *ProductSimpleDTO `json:"product"`
}

type InventoryResponseDTO struct {
	ID            uint      `json:"id"`
	ProductID     uint      `json:"product_id"`
	StoreID       uuid.UUID `json:"store_id"`
	Quantity      int       `json:"quantity"`
	MinStockLevel int       `json:"min_stock_level"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
