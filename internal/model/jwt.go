package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTData struct {
	jwt.RegisteredClaims

	UserID uuid.UUID `json:"userId"`
}
