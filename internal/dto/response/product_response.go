package response

import (
	"time"

	"github.com/google/uuid"
)

type ProductWithCategoryResponseDTO struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	SKU         string             `json:"sku"`
	ImageURL    string             `json:"image_url"`
	Size        string             `json:"size"`
	Color       string             `json:"color"`
	Price       float64            `json:"price"`
	Category    *CategorySimpleDTO `json:"category"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type ProductWithInventoryResponse struct {
	Product   *ProductResponseDTO  `json:"product"`
	Inventory *InventoryProductDTO `json:"inventory"`
}

type ProductResponseDTO struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	SKU         string             `json:"sku"`
	ImageURL    string             `json:"image_url"`
	Size        string             `json:"size"`
	Color       string             `json:"color"`
	Price       float64            `json:"price"`
	Category    *CategorySimpleDTO `json:"category"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type InventoryProductDTO struct {
	ID            uint      `json:"id"`
	StoreID       uuid.UUID `json:"store_id"`
	Quantity      int       `json:"quantity"`
	MinStockLevel int       `json:"min_stock_level"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
