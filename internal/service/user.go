package service

import (
	"context"
	"net/http"

	"github.com/auto-hh/backend/internal/model"
	"github.com/auto-hh/backend/internal/repository"
	"github.com/google/uuid"
)

type User struct {
	repository repository.IProfile
	client     *http.Client
}

func NewUser(repository repository.IProfile, client *http.Client) *User {
	return &User{
		repository: repository,
		client:     client,
	}
}

func (u *User) GetUserInfo(ctx context.Context, userID uuid.UUID) (model.Profile, error) {
	return u.repository.GetProfileData(ctx, userID)
}

func (u *User) IsProfileExistsByUserID(ctx context.Context, userID uuid.UUID) (bool, error) {
	return u.repository.IsProfileExistsByUserID(ctx, userID)
}

func (u *User) UpdateUserInfo(ctx context.Context, userID uuid.UUID, profile model.Profile) error {
	return u.repository.InsertOrUpdate(ctx, userID, profile)
}
