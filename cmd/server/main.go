package main

import (
	"context"

	"github.com/auto-hh/backend/config"
	"github.com/auto-hh/backend/internal/app"
	"github.com/auto-hh/backend/pkg/postgres"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := postgres.NewPool(ctx, config.PostgresDSN())
	if err != nil {
		panic(err)
	}

	server, err := app.InitServer(pool, config.SecretKey(), config.LLMPath())
	if err != nil {
		panic(err)
	}

	err = server.Start(config.Address())
	if err != nil {
		panic(err)
	}
}
