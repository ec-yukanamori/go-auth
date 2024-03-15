package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppEnv AppEnv `envconfig:"APP_ENV" default:"development"`

	Server struct {
		Port string `envconfig:"APP_SERVER_PORT" default:"1323"`
	}

	SecretKey string `envconfig:"SECRET_KEY" default:"my_secret_key"`
}

var config *Config

func init() {
	c := &Config{}
	if err := envconfig.Process("myapp", c); err != nil {
		panic(fmt.Errorf("error loading config from env: %s", err))
	}

	config = c
}
