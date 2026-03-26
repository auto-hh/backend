package service

import (
	"context"

	"github.com/auto-hh/backend/internal/model"
	"github.com/google/uuid"
)

type IAuth any

type ILLM interface {
	FindVacancies(ctx context.Context, userID uuid.UUID) ([]model.Vacancy, error)
	Analysis(ctx context.Context, userID uuid.UUID) ([]model.Attribute, error)
}
