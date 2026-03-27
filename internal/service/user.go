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

func (user *User) GetUserInfo(ctx context.Context, userID uuid.UUID) (model.Profile, error) {
	rawUserInfo, err := user.repository.GetProfileData(ctx, userID)
	if err != nil {
		return model.Profile{}, err
	}

	return rawUserInfo, nil
}

func (user *User) IsProfileExistsByUserID(ctx context.Context, userID uuid.UUID) (bool, error) {
	return user.repository.IsProfileExistsByUserID(ctx, userID)
}
