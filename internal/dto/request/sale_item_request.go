package request

type CreateSaleItemRequestDTO struct {
	SaleID     int     `json:"sale_id" validate:"required"`
	ProductID  int     `json:"product_id" validate:"required"`
	Quantity   int     `json:"quantity" validate:"required,min=1"`
	UnitPrice  float64 `json:"unit_price" validate:"required,gt=0"`
	TotalPrice float64 `json:"total_price" validate:"required,gt=0"`
}
