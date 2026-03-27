package app

import (
	"net/http"

	"github.com/auto-hh/backend/internal/service"
)

type Services struct {
	auth service.IAuth
	llm  service.ILLM
}

func InitServices(repositories *Repositories, llmPath string) *Services {
	client := &http.Client{}

	return &Services{
		auth: service.NewAuth(repositories.user),
		llm:  service.NewLLM(repositories.profile, client, llmPath),
	}
}
