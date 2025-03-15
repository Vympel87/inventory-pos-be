package request

type CreateProductRequest struct {
	CategoryID    *int    `json:"category_id" validate:"omitempty,number"`
	Name          string  `json:"name" validate:"required"`
	Description   string  `json:"description" validate:"omitempty"`
	SKU           string  `json:"sku" validate:"required"`
	ImageURL      string  `json:"image_url" validate:"omitempty,url"`
	Size          string  `json:"size" validate:"omitempty,max=5"`
	Color         string  `json:"color" validate:"omitempty,max=30"`
	Price         float64 `json:"price" validate:"required,gt=0"`
	StoreID       string  `json:"store_id" validate:"required,uuid4"`
	Quantity      int     `json:"quantity" validate:"required,gte=0"`
	MinStockLevel int     `json:"min_stock_level" validate:"required,gte=0"`
}

type UpdateProductRequest struct {
	CategoryID  *int    `json:"category_id" validate:"omitempty,number"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"omitempty"`
	SKU         string  `json:"sku" validate:"required"`
	ImageURL    string  `json:"image_url" validate:"omitempty,url"`
	Size        string  `json:"size" validate:"omitempty,max=5"`
	Color       string  `json:"color" validate:"omitempty,max=30"`
	Price       float64 `json:"price" validate:"required,gt=0"`
}
