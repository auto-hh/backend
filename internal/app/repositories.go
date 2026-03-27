package app

import (
	"github.com/auto-hh/backend/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	txManager repository.TransactionManager
	user      repository.IUser
	profile   repository.IProfile
}

func InitRepositories(pool *pgxpool.Pool) *Repositories {
	executor := repository.NewExecutor(pool)

	return &Repositories{
		txManager: executor,
		user:      repository.NewUser(executor),
		profile:   repository.NewProfile(executor),
	}
}
