package response

import (
	"time"
)

type SaleItemResponseDTO struct {
	ID         int               `json:"id"`
	SaleID     int               `json:"sale_id"`
	ProductID  int               `json:"product_id"`
	Quantity   int               `json:"quantity"`
	UnitPrice  float64           `json:"unit_price"`
	TotalPrice float64           `json:"total_price"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	Product    *ProductSimpleDTO `json:"product,omitempty"`
	Sale       *SaleSimpleDTO    `json:"sale,omitempty"`
}

type CompleteSaleItemResponseDTO struct {
	ID         int       `json:"id"`
	SaleID     int       `json:"sale_id"`
	ProductID  int       `json:"product_id"`
	Quantity   int       `json:"quantity"`
	UnitPrice  float64   `json:"unit_price"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
