package main

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func getBearerTokenFromHeader(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthorizationHeader
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return "", ErrInvalidAuthorizationHeader
	}

	return authParts[1], nil
}
