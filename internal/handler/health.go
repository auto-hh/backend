package handler

import "github.com/labstack/echo/v5"

type Health struct {
}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) Health(ctx *echo.Context) error {
	return nil
}
