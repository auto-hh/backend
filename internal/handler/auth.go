package handler

import (
	"net/http"
	"time"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/model"
	"github.com/auto-hh/backend/internal/service"
	"github.com/labstack/echo/v5"
)

type Auth struct {
	serviceAuth service.IAuth
	stateExpirationDuration time.Duration
	jwtExpirationDuration time.Duration
	siteURL string
}

func NewAuth(serviceAuth service.IAuth, stateExpirationDuration, jwtExpirationDuration time.Duration, siteURL string) *Auth {
	return &Auth{
		serviceAuth: serviceAuth,
		stateExpirationDuration: stateExpirationDuration,
		jwtExpirationDuration: jwtExpirationDuration,
		siteURL: siteURL,
	}
}

// @Tags         auth
// @Success      200
// @Router       /auth/begin [get]
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

// @Tags         auth
// @Success      200
// @Router       /auth/complete [get]
func (a *Auth) Complete(ctx *echo.Context) error {
	stateCookie, err := ctx.Cookie(domain.CookieState)
	if err != nil {
		err = domain.NewBadRequest(domain.CodeBadRequest, "state cookie error", err)
		return domain.MapAppError(ctx, err)
	}

	var complete model.Complete
	err = ctx.Bind(&complete)
	if err != nil {
		err = domain.NewBadRequest(domain.CodeBadRequest, "query params are invalid", err)
		return domain.MapAppError(ctx, err)
	}
	authJWTToken, err := a.serviceAuth.Complete(ctx.Request().Context(), stateCookie.Value, complete)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}
	authCookie := &http.Cookie{
		Name: domain.CookieAuthJWT,
		Value: authJWTToken,
		Path: "/",
		Expires: time.Now().Add(a.jwtExpirationDuration),
		MaxAge: int(a.jwtExpirationDuration.Seconds()),
		Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	ctx.SetCookie(authCookie)
	return ctx.Redirect(http.StatusFound, a.siteURL)
}
