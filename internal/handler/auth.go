package handler

import (
	"net/http"
	"time"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/service"
	"github.com/labstack/echo/v5"
)

type Auth struct {
	serviceAuth service.IAuth
	stateExpirationDuration time.Duration
}

func NewAuth(serviceAuth service.IAuth, stateExpirationDuration time.Duration) *Auth {
	return &Auth{
		serviceAuth: serviceAuth,
		stateExpirationDuration: stateExpirationDuration,
	}
}

func (a *Auth) Begin(ctx *echo.Context) error {
	stateJWTToken, redirectURL, err := a.serviceAuth.Begin()
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	stateCookie := &http.Cookie{
		Name: domain.CookieState,
		Value: stateJWTToken,
		Path: "/auth/complete",
		Expires: time.Now().Add(a.stateExpirationDuration),
		MaxAge: int(a.stateExpirationDuration.Seconds()),
		Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	ctx.SetCookie(stateCookie)

	return ctx.Redirect(http.StatusFound, redirectURL.String())
}

func (a *Auth) Complete(_ *echo.Context) error {
	return nil
}
