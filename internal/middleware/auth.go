package middleware

import (
	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

const KeyToken string = "user"

func GetJWTData(ctx *echo.Context) (*model.JWTData, error) {
	token, err := echo.ContextGet[*jwt.Token](ctx, KeyToken)
	if err != nil {
		return nil, domain.NewUnauthorized(
			domain.CodeUnauthorized,
			"token key not found in context",
			err,
		)
	}

	jwtData, ok := token.Claims.(*model.JWTData)
	if !ok {
		return nil, domain.NewForbidden(domain.CodeForbidden, "invalid token claims")
	}

	return jwtData, nil
}

func GetUserID(ctx *echo.Context) (uuid.UUID, error) {
	jwtData, err := GetJWTData(ctx)
	if err != nil {
		return uuid.UUID{}, err
	}

	return jwtData.UserID, nil
}
