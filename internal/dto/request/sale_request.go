package request

import (
	"time"

	"github.com/google/uuid"
)

type CreateSaleItem struct {
	ProductID  int     `json:"product_id" validate:"required"`
	Quantity   int     `json:"quantity" validate:"required,min=1"`
	UnitPrice  float64 `json:"unit_price" validate:"required"`
	TotalPrice float64 `json:"total_price" validate:"required"`
}

type CreateSaleRequest struct {
	StoreID       uuid.UUID         `json:"store_id" validate:"required"`
	UserID        *uuid.UUID        `json:"user_id,omitempty"`
	SaleDate      time.Time         `json:"sale_date" validate:"required"`
	SaleNumber    string            `json:"sale_number" validate:"required"`
	PaymentMethod string            `json:"payment_method" validate:"required"`
	TotalAmount   float64           `json:"total_amount" validate:"required"`
	Items         []*CreateSaleItem `json:"items" validate:"required,dive"`
}

type UpdateSaleRequest struct {
	StoreID       uuid.UUID  `json:"store_id"`
	UserID        *uuid.UUID `json:"user_id,omitempty"`
	SaleDate      time.Time  `json:"sale_date"`
	SaleNumber    string     `json:"sale_number"`
	PaymentMethod string     `json:"payment_method"`
	TotalAmount   float64    `json:"total_amount"`
}
