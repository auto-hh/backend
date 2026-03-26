package service

import "github.com/auto-hh/backend/internal/repository"

type Auth struct {
	repositoryUser repository.IUser
}

func NewAuth(repositoryUser repository.IUser) *Auth {
	return &Auth{
		repositoryUser: repositoryUser,
	}
}
