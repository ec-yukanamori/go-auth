package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Server struct {
		Port string `envconfig:"APP_SERVER_PORT" default:"1323"`
	}

	RDS struct {
		Host     string `envconfig:"APP_RDS_HOST" default:"redis"`
		Port     string `envconfig:"APP_RDS_PORT" default:"6379"`
		Password string `envconfig:"APP_RDS_PASSWORD" default:""`
		DB       int    `envconfig:"APP_RDS_DB" default:"0"`
	}

	JWT struct {
		ExpirationMinutes int `envconfig:"APP_JWT_EXPIRATION_MINUTES" default:"60"`
	}
}

var cfg *config

func init() {
	cfg = &config{}
	if err := envconfig.Process("myapp", cfg); err != nil {
		panic(fmt.Errorf("failed to load config from env: %s", err))
	}
}
