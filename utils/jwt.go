package utils

import (
	"fmt"
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

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signin method")
		}
		return []byte(secretKet), nil
	})

	if err != nil {
		return 0, err
	}

	if !parsedToken.Valid {
		return 0, fmt.Errorf("token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("Invalid token")
	}

	userId := int64(claims["userId"].(float64))
	return userId, nil
}
