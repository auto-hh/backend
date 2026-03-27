package app

import (
	"fmt"

	"github.com/auto-hh/backend/config"
	"github.com/auto-hh/backend/internal/domain"
	"github.com/auto-hh/backend/internal/middleware"
	"github.com/auto-hh/backend/internal/model"
	"github.com/auto-hh/backend/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	echomw "github.com/labstack/echo/v5/middleware"
)

func InitServer(config *config.Config, pool *pgxpool.Pool) (*echo.Echo, error) {
	server := echo.New()
	server.Logger = logger.InitLogger(config.LogLevel)

	repositories := InitRepositories(pool)
	services := InitServices(config, repositories)
	handlers := InitHandlers(config, services)

	AddHandlers(config, server, handlers)

	return server, nil
}

func AddHandlers(config *config.Config, server *echo.Echo, handlers *Handlers) {
	jwtConfig := InitJWTConfig(config.SecretKey)

	server.Use(echomw.Recover())
	server.Use(echomw.RequestLogger())

	server.GET("/health", handlers.health.Health)

	groupAuth := server.Group("/auth")
	groupUser := server.Group("/user", echojwt.WithConfig(jwtConfig))
	groupLLM := server.Group("/llm", echojwt.WithConfig(jwtConfig))

	groupAuth.GET("/begin", handlers.auth.Begin)
	groupAuth.GET("/complete", handlers.auth.Complete)

	groupUser.GET("/me", handlers.user.Me)
	groupUser.GET("/has-profile", handlers.user.HasProfile)
	groupUser.GET("/profile", handlers.user.Profile)

	groupLLM.POST("/vacancies", handlers.llm.FindVacancies)
	groupLLM.POST("/analysis", handlers.llm.Analysis)
	groupLLM.POST("/generate", handlers.llm.Analysis)
}

func InitJWTConfig(secretKey []byte) echojwt.Config {
	//nolint:gosec
	return echojwt.Config{
		ErrorHandler: func(ctx *echo.Context, err error) error {
			return domain.MapAppError(
				ctx,
				domain.NewUnauthorized(domain.CodeUnauthorized, "jwt middleware error", err),
			)
		},
		SigningKey:  secretKey,
		ContextKey:  middleware.KeyToken,
		TokenLookup: fmt.Sprintf("cookie:%s", domain.CookieAuthJWT),
		NewClaimsFunc: func(_ *echo.Context) jwt.Claims {
			return new(model.JWTAuthData)
		},
	}
}
