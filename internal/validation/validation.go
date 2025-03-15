package validation

import "github.com/Vympel87/inventory-pos-be/internal/utils"

var validator = utils.NewValidator()

func Validate(data interface{}) error {
	return validator.ValidateStruct(data)
}
