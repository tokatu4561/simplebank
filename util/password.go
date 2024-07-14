package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword (password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed, to hashpassword: %w", err)
	}

	return string(hashPassword), nil
}

func CheckPassword (hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}