package app

import (
	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/handler"
	"github.com/auto-hh/backend/internal/middleware"
	"github.com/auto-hh/backend/internal/model"
	"github.com/auto-hh/backend/internal/repository"
	"github.com/auto-hh/backend/internal/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	echomw "github.com/labstack/echo/v5/middleware"
)

type Handlers struct {
	auth handler.Auth
}

type Services struct {
	auth service.IAuth
}

type Repositories struct {
	txManager repository.TransactionManager
	user      repository.IUser
}

func InitServer(pool *pgxpool.Pool, secretKey []byte) (*echo.Echo, error) {
	server := echo.New()

	repositories := InitRepositories(pool)
	services := InitServices(repositories)
	handlers := InitHandlers(services)

	AddHandlers(server, handlers, secretKey)

	return server, nil
}

func AddHandlers(server *echo.Echo, handlers *Handlers, secretKey []byte) {
	jwtConfig := InitJWTConfig(secretKey)

	server.Use(echomw.Recover())
	server.Use(echomw.RequestLogger())
	//TODO: add CORS middleware

	groupAuth := server.Group("/auth")
	groupLLM := server.Group("/llm", echojwt.WithConfig(jwtConfig))

	groupAuth.GET("/begin", handlers.auth.Begin)
	groupAuth.GET("/complete", handlers.auth.Complete)

	groupLLM.POST("/search", ...)
	groupLLM.POST("/analyze", ...)
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
		user:      repository.NewUser(executor),
	}
}

func InitJWTConfig(secretKey []byte) echojwt.Config {
	//nolint:gosec
	return echojwt.Config{
		ErrorHandler: func(ctx *echo.Context, err error) error {
			return domain.MapAppError(
				ctx,
				domain.NewUnauthorized(domain.CodeUnauthorized, "handled jwt error", err),
			)
		},
		SigningKey:  secretKey,
		ContextKey:  middleware.KeyToken,
		TokenLookup: "cookie:auto-hh-access-token",
		NewClaimsFunc: func(_ *echo.Context) jwt.Claims {
			return new(model.JWTData)
		},
	}
}
