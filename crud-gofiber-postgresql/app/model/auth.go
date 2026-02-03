package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Nip      string `json:"nip"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

// JWTClaims represents the claims in a JWT token
type JWTClaims struct {
	UserID int    `json:"user_id"`
	Nip    string `json:"nip"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
