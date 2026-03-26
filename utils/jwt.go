package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKet = "secret"

func GenerateToken(userId int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Hour).Unix(),
	})

	return token.SignedString([]byte(secretKet))
}
