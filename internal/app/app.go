package app

import (
	"net/http"

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
	auth *handler.Auth
	llm  *handler.LLM
}

type Services struct {
	auth service.IAuth
	llm  service.ILLM
}

type Repositories struct {
	txManager repository.TransactionManager
	user      repository.IUser
	profile   repository.IProfile
}

func InitServer(pool *pgxpool.Pool, secretKey []byte, llmPath string) (*echo.Echo, error) {
	server := echo.New()

	repositories := InitRepositories(pool)
	services := InitServices(repositories, llmPath)
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

	groupLLM.POST("/vacancies", handlers.llm.FindVacancies)
	groupLLM.POST("/analysis", handlers.llm.Analysis)
}

func InitHandlers(services *Services) *Handlers {
	return &Handlers{
		auth: handler.NewAuth(services.auth),
		llm:  handler.NewLLM(services.llm),
	}
}

func InitServices(repositories *Repositories, llmPath string) *Services {
	client := &http.Client{}
	return &Services{
		auth: service.NewAuth(repositories.user),
		llm:  service.NewLLM(repositories.profile, client, llmPath),
	}
}

func InitRepositories(pool *pgxpool.Pool) *Repositories {
	executor := repository.NewExecutor(pool)

	return &Repositories{
		txManager: executor,
		user:      repository.NewUser(executor),
		profile:   repository.NewProfile(executor),
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
