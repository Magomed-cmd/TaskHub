package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GetHash(password string) (string, error) {

	if password == "" {
		return "", errors.New("password is empty")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	hash := string(bytes)
	return hash, err
}

func CheckPassword(plainPassword string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
