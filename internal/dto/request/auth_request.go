package request

type RegisterRequestDTO struct {
	StoreName    string `json:"store_name" validate:"required"`
	StoreAddress string `json:"store_address" validate:"required"`
	StorePhone   string `json:"store_phone" validate:"required"`
	StoreEmail   string `json:"store_email" validate:"required,email"`
	UserName     string `json:"user_name" validate:"required"`
	UserEmail    string `json:"user_email" validate:"required,email"`
	UserPassword string `json:"user_password" validate:"required,min=6"`
}

type LoginRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
