package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}
	defer l.Sync()

	logger = l
}

var requestLoggerConfig = middleware.RequestLoggerConfig{
	LogRequestID: true,
	LogRemoteIP:  true,
	LogHost:      true,
	LogMethod:    true,
	LogURI:       true,
	LogUserAgent: true,
	LogStatus:    true,
	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
		logger.Info("request",
			zap.String("id", v.RequestID),
			zap.String("remote_ip", v.RemoteIP),
			zap.String("host", v.Host),
			zap.String("method", v.Method),
			zap.String("uri", v.URI),
			zap.String("user_agent", v.UserAgent),
			zap.Int("status", v.Status),
		)
		return nil
	},
}
