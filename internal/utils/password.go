package utils

import (
	"golang.org/x/crypto/bcrypt"

	enum "github.com/Vympel87/inventory-pos-be/internal/constants"
	"github.com/Vympel87/inventory-pos-be/internal/exception"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", exception.NewInternalException("Failed to hash password", err.Error(), enum.INTERNAL_EXCEPTION)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
