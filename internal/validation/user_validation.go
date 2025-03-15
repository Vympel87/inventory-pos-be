package validation

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,oneof=OWNER SUPERVISOR ADMIN"`
	StoreID  string `json:"storeId" validate:"required,uuid4"`
}

func ValidateCreateUserRequest(data *CreateUserRequest) error {
	return Validate(data)
}
