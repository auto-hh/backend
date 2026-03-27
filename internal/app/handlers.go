package app

import "github.com/auto-hh/backend/internal/handler"

type Handlers struct {
	health *handler.Health
	auth *handler.Auth
	llm  *handler.LLM
}

func InitHandlers(services *Services) *Handlers {
	return &Handlers{
		health: handler.NewHealth(),
		auth: handler.NewAuth(services.auth),
		llm:  handler.NewLLM(services.llm),
	}
}
