package app

import (
	"net/http"

	"github.com/auto-hh/backend/config"
	"github.com/auto-hh/backend/internal/service"
)

type Services struct {
	auth service.IAuth
	llm  service.ILLM
}

func InitServices(config *config.Config, repositories *Repositories) *Services {
	client := &http.Client{}

	return &Services{
		auth: service.NewAuth(repositories.user, config.SecretKey, config.ClientID, config.RedirectURI),
		llm:  service.NewLLM(repositories.profile, client, config.LLMPath),
	}
}
