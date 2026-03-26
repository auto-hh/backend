package handler

import "github.com/auto-hh/backend/internal/service"

type Auth struct {
	serviceAuth service.IAuth
}

func NewAuth(serviceAuth service.IAuth) *Auth {
	return &Auth{
		serviceAuth: serviceAuth,
	}
}
