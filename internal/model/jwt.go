package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTAuthData struct {
	jwt.RegisteredClaims

	UserID uuid.UUID `json:"userId"`
}

type JWTStateData struct {
	jwt.RegisteredClaims

	State string `json:"state"`
}
