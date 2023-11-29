package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var ErrWrongPassword = errors.New("wrong password")

func BcryptPassword(password string) (string, error) {
	const DefaultCost = 10

	HashPassword, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	return string(HashPassword), err
}

func ValidatePassword(password string, encryptedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	if err != nil {
		return false, ErrWrongPassword
	}
	return true, nil
}
