package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	Uuid  string `json:"uuid"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID   string `json:"id"`
	Uuid string `json:"uuid"`
	jwt.RegisteredClaims
}
