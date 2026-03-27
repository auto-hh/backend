package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type User struct {
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
	return ctx.NoContent(http.StatusNoContent)
}
