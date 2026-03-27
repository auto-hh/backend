package service

import (
	"context"
	"net/url"

	"github.com/auto-hh/backend/internal/model"
	"github.com/google/uuid"
)

type IAuth interface {
	Begin() (string, *url.URL, error)
	Complete(
		ctx context.Context,
		stateJWTToken string,
		complete model.Complete,
	) (authJWTToken string, err error)
}

type ILLM interface {
	FindVacancies(ctx context.Context, userID uuid.UUID) ([]model.Vacancy, error)
	Analysis(ctx context.Context, userID uuid.UUID) ([]model.Attribute, error)
	GetCoverLetter(ctx context.Context, userID uuid.UUID, vacancy model.Vacancy) (string, error)
}

type IUser interface {
	GetUserInfo(ctx context.Context, userID uuid.UUID) (model.Profile, error)
	IsProfileExistsByUserID(ctx context.Context, userID uuid.UUID) (bool, error)
	UpdateUserInfo(ctx context.Context, userID uuid.UUID, profile model.Profile) error
}
