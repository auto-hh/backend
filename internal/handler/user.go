package handler

import (
	"net/http"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/middleware"
	"github.com/auto-hh/backend/internal/service"
	"github.com/labstack/echo/v5"
)

type User struct {
	service service.IUser
}

func NewUser() *User {
	return &User{}
}

func (u *User) Me(ctx *echo.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (u *User) HasProfile(ctx *echo.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (u *User) Profile(ctx *echo.Context) error {

	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	profileInfo, err := u.service.GetUserInfo(ctx.Request().Context(), userID)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, profileInfo)
}
