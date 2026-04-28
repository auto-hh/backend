package app

import (
	"net/http"

	"github.com/auto-hh/backend/config"
	"github.com/auto-hh/backend/internal/service"
)

type Services struct {
	auth service.IAuth
	user service.IUser
	llm  service.ILLM
}

func InitServices(config *config.Config, repositories *Repositories) *Services {
	client := &http.Client{}

	return &Services{
		auth: service.NewAuth(
			repositories.user,
			client,
			config.SecretKey,
			config.ClientID,
			config.ClientSecret,
			config.RedirectURI,
			config.AppName,
			config.AppVersion,
			config.DevContact,
		),
		user: service.NewUser(repositories.profile, client),
		llm:  service.NewLLM(repositories.profile, client, config.LLMPath),
	}
}
