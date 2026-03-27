package app

import (
	"github.com/auto-hh/backend/config"
	"github.com/auto-hh/backend/internal/handler"
)

type Handlers struct {
	health *handler.Health
	auth *handler.Auth
	user *handler.User
	llm  *handler.LLM
}

func InitHandlers(config *config.Config, services *Services) *Handlers {
	return &Handlers{
		health: handler.NewHealth(),
		auth: handler.NewAuth(services.auth, config.StateExpirationDuration, config.JWTExpirationDuration, config.SiteURL),
		user: handler.NewUser(),
		llm:  handler.NewLLM(services.llm),
	}
}
