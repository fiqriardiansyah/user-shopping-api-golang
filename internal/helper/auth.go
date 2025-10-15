package helper

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTClaims struct {
	Email  string    `json:"email"`
	UserId uuid.UUID `json:"user_id"`
	Roles  []string  `json:"roles"`
	jwt.RegisteredClaims
}

type GenerateTokenParam struct {
	UserId   uuid.UUID
	Email    string
	Roles    []string
	Duration time.Time
	Secret   string
}

func GenerateToken(param GenerateTokenParam) (string, error) {
	claims := &JWTClaims{
		Email:  param.Email,
		UserId: param.UserId,
		Roles:  param.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(param.Duration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(param.Secret))
}

func ValidateToken(tokenStr string, secret string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func ValidRedirectUrl(redirectUri string) bool {
	orderService := os.Getenv("ORDER_SERVICE_URL")
	services := []string{
		orderService,
	}

	for _, s := range services {
		if strings.HasPrefix(redirectUri, s) {
			return true
		}
	}

	return false
}
