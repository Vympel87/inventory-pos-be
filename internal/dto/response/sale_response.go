package response

import (
	"time"

	"github.com/google/uuid"
)

type SaleResponse struct {
	ID            int        `json:"id"`
	StoreID       uuid.UUID  `json:"store_id"`
	UserID        *uuid.UUID `json:"user_id,omitempty"`
	SaleDate      time.Time  `json:"sale_date"`
	SaleNumber    string     `json:"sale_number"`
	PaymentMethod string     `json:"payment_method"`
	TotalAmount   float64    `json:"total_amount"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type SaleData struct {
	ID            int        `json:"id"`
	StoreID       uuid.UUID  `json:"store_id"`
	UserID        *uuid.UUID `json:"user_id,omitempty"`
	SaleDate      time.Time  `json:"sale_date"`
	SaleNumber    string     `json:"sale_number"`
	PaymentMethod string     `json:"payment_method"`
	TotalAmount   float64    `json:"total_amount"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type SaleItemData struct {
	ID         int       `json:"id"`
	SaleID     int       `json:"sale_id"`
	ProductID  int       `json:"product_id"`
	Quantity   int       `json:"quantity"`
	UnitPrice  float64   `json:"unit_price"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateSaleResponse struct {
	Sale      *SaleData       `json:"sale"`
	SaleItems []*SaleItemData `json:"sale_items"`
}
