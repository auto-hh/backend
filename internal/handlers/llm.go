package handlers

import (
	"net/http"

	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/middleware"
	"github.com/auto-hh/backend/internal/service"
	"github.com/labstack/echo/v5"
)

type LLM struct {
	service service.ILLM
}

func NewLLM(service service.ILLM) *LLM {
	return &LLM{
		service: service,
	}
}

func (llm *LLM) FindVacancies(ctx *echo.Context) error {
	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	vacancies, err := llm.service.FindVacancies(ctx.Request().Context(), userID)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	return domain.JSON(ctx, http.StatusOK, vacancies)
}

func (llm *LLM) Analysis(ctx *echo.Context) error {
	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	scores, err := llm.service.Analysis(ctx.Request().Context(), userID)
	if err != nil {
		return domain.MapAppError(ctx, err)
	}

	return domain.JSON(ctx, http.StatusOK, scores)
}
