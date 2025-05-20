package utils

import (
	"TaskHub/internal/pkg/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type CustomClaims struct {
	Email string
	UUID  uuid.UUID
	jwt.RegisteredClaims
}

func GenerateJWT(user model.User, secret string) (*string, error) {
	claims := &CustomClaims{
		Email: user.Email,
		UUID:  user.UUID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}
