package service

import (
	"context"
	"net/url"

	"github.com/auto-hh/backend/internal/model"
	"github.com/google/uuid"
)

type IAuth interface {
	Begin() (string, *url.URL, error)
}

type ILLM interface {
	FindVacancies(ctx context.Context, userID uuid.UUID) ([]model.Vacancy, error)
	Analysis(ctx context.Context, userID uuid.UUID) ([]model.Attribute, error)
}
