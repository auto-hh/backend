package main

import (
	"context"

	"github.com/auto-hh/backend/config"
	"github.com/auto-hh/backend/internal/app"
	"github.com/auto-hh/backend/pkg/postgres"
	_ "github.com/auto-hh/backend/docs"
)

// @title Auto HH
// @version 0.1.0
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
