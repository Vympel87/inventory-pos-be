package response

import (
	"github.com/google/uuid"

	"time"
)

type StoreSimpleDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	StoreType string    `json:"store_type"`
}

type CategorySimpleDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductSimpleDTO struct {
	ID       int                `json:"id"`
	Name     string             `json:"name"`
	ImageURL string             `json:"image_url"`
	Size     string             `json:"size"`
	Color    string             `json:"color"`
	Price    float64            `json:"price"`
	Category *CategorySimpleDTO `json:"category"`
}

type SaleSimpleDTO struct {
	ID            int       `json:"id"`
	StoreID       uuid.UUID `json:"store_id"`
	UserID        uuid.UUID `json:"user_id,omitempty"`
	SaleDate      time.Time `json:"sale_date"`
	SaleNumber    string    `json:"sale_number"`
	PaymentMethod string    `json:"payment_method"`
	TotalAmount   float64   `json:"total_amount"`
}
