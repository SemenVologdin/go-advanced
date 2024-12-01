package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type App struct {
	Server  Server
	Mail    Mail
	Storage Storage
}

func New() (App, error) {
	if err := godotenv.Load(); err != nil {
		return App{}, err
	}

	var app App
	if err := env.Parse(&app); err != nil {
		return App{}, err
	}

	return app, nil
}
