package utils

import "golang.org/x/crypto/bcrypt"

func BcryptPassWord(password string) (string, error) {
	const DefaultCost = 10

	HashPassWord, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	return string(HashPassWord), err
}

func CheckPassWord(password string, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err
}
