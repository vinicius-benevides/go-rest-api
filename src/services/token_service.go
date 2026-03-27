package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateToken(userID int64, email string) (string, error)
	VerifyToken(token string) (int64, error)
}

type jwtService struct {
	secret     string
	expiration time.Duration
}

func NewTokenService(secret string, expiration time.Duration) TokenService {
	return &jwtService{
		secret:     secret,
		expiration: expiration,
	}
}

func (s *jwtService) GenerateToken(userID int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userID,
		"email":  email,
		"exp":    time.Now().Add(s.expiration).Unix(),
	})

	return token.SignedString([]byte(s.secret))
}

func (s *jwtService) VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(s.secret), nil
	})

	if err != nil {
		return 0, err
	}

	if !parsedToken.Valid {
		return 0, fmt.Errorf("token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["userId"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid token payload")
	}

	return int64(userID), nil
}
