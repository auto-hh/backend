package main

import (
	"context"

	"github.com/auto-hh/backend/config"
	_ "github.com/auto-hh/backend/docs"
	"github.com/auto-hh/backend/internal/app"
	"github.com/auto-hh/backend/pkg/postgres"
)

// @title						Auto HH
// @version					0.1.0
// @securityDefinitions.apikey	BearerAuth
// @in							cookie
// @name						auto-hh-access-key.
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

	server, err := app.InitServer(config, pool)
	if err != nil {
		panic(err)
	}

	err = server.Start(config.Address())
	if err != nil {
		panic(err)
	}
}
