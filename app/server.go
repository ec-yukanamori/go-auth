package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	echo *echo.Echo
}

var srv *server

func init() {
	srv = &server{
		echo: echo.New(),
	}

	srv.echo.Use(middleware.Recover())
	srv.echo.Use(middleware.RequestID())
	srv.echo.Use(middleware.RequestLoggerWithConfig(requestLoggerConfig))

	srv.echo.Validator = newValidator()

	srv.echo.GET("/ping", ping)

	tokenHandler := assembleTokenHandler()
	srv.echo.POST("/token", tokenHandler.generate)
	srv.echo.GET("/token", tokenHandler.verify)
	srv.echo.PUT("/token", tokenHandler.refresh)
	srv.echo.DELETE("/token", tokenHandler.delete)
}

func Start() {
	if err := srv.echo.Start(fmt.Sprintf(":%s", cfg.Server.Port)); err != nil {
		panic(fmt.Sprintf("failed to start server: %v", err))
	}
}
