package handlers

import (
	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/service"
	"github.com/labstack/echo/v5"
)

type LLM struct {
	service service.ILLM
}

func (llm *LLM) FindVacancies(ctx *echo.Context) error {
	err := llm.service.FindVacancies()
	if err != nil {
		return domain.MapAppError(ctx, err)
	}
	return nil
}
