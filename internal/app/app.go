package app

import (
	"context"

	"github.com/auto-hh/backend/config"
	"github.com/auto-hh/backend/internal/handler"
	"github.com/auto-hh/backend/internal/repository"
	"github.com/auto-hh/backend/internal/service"
	"github.com/auto-hh/backend/pkg/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
)

type Handlers struct {
	auth handler.Auth
}

type Services struct {
	auth service.IAuth
}

type Repositories struct {
	txManager repository.TransactionManager
	user repository.IUser
}

func InitServer() (*echo.Echo, error) {
	server := echo.New()

	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := postgres.NewPool(ctx, config.PostgresDSN())

	repositories := InitRepositories(pool)
	services := InitServices(repositories)
	handlers := InitHandlers(services)

	AddHandlers(server, handlers)

	return server, nil
}

func AddHandlers(server *echo.Echo, handlers *Handlers) {
	
}

func InitHandlers(services *Services) *Handlers {
	return &Handlers{
		auth: *handler.NewAuth(services.auth),
	}
}

func InitServices(repositories *Repositories) *Services {
	return &Services{
		auth: service.NewAuth(repositories.user),
	}
}

func InitRepositories(pool *pgxpool.Pool) *Repositories {
	executor := repository.NewExecutor(pool)
	return &Repositories{
		txManager: executor,
		user: repository.NewUser(executor),
	}
}
