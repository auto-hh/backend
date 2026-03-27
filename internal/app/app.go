package app

import (
	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/middleware"
	"github.com/auto-hh/backend/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	echomw "github.com/labstack/echo/v5/middleware"
)


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

	server.GET("/health", handlers.health.Health)

	groupAuth := server.Group("/auth")
	groupLLM := server.Group("/llm", echojwt.WithConfig(jwtConfig))

	groupAuth.GET("/begin", handlers.auth.Begin)
	groupAuth.GET("/complete", handlers.auth.Complete)

	groupLLM.POST("/vacancies", handlers.llm.FindVacancies)
	groupLLM.POST("/analysis", handlers.llm.Analysis)
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
