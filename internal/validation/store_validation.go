package validation

type CreateStoreRequest struct {
	Name      string `json:"name" validate:"required,max=255"`
	Address   string `json:"address" validate:"omitempty"`
	Phone     string `json:"phone" validate:"omitempty,max=50"`
	Email     string `json:"email" validate:"omitempty,email,max=255"`
	StoreType string `json:"storeType" validate:"required,oneof=CENTRAL BRANCH"`
	ParentID  string `json:"parentStoreId" validate:"omitempty,uuid4"`
}

func ValidateCreateStoreRequest(data *CreateStoreRequest) error {
	return Validate(data)
}
