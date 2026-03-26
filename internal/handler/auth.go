package handler

import (
	"github.com/auto-hh/backend/internal/service"
	"github.com/labstack/echo/v5"
)

type Auth struct {
	serviceAuth service.IAuth
}

func NewAuth(serviceAuth service.IAuth) *Auth {
	return &Auth{
		serviceAuth: serviceAuth,
	}
}

func (a *Auth) Begin(_ *echo.Context) error {
	return nil
}

func (a *Auth) Complete(_ *echo.Context) error {
	return nil
}
