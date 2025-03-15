package request

import (
	"github.com/google/uuid"
)

type CreateInventoryRequestDTO struct {
	ProductID     uint      `json:"product_id" binding:"required"`
	StoreID       uuid.UUID `json:"store_id" binding:"required"`
	Quantity      int       `json:"quantity" binding:"required,min=0"`
	MinStockLevel int       `json:"min_stock_level" binding:"required,min=0"`
}
