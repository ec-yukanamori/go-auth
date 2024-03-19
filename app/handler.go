package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var ping = func(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

type tokenHandler struct {
	usecase tokenUsecase
}

func newTokenHandler(usecase tokenUsecase) *tokenHandler {
	return &tokenHandler{usecase: usecase}
}

type (
	generateTokenResponse struct {
		token string
	}
)

func (h *tokenHandler) generate(c echo.Context) error {
	var req generateTokenInput
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(req); err != nil {
		return echo.ErrBadRequest
	}

	token, err := h.usecase.generate(req)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, generateTokenResponse{token: token})
}

func (h *tokenHandler) verify(c echo.Context) error {
	token, err := getBearerTokenFromHeader(c)
	if err != nil {
		if err == ErrNoAuthorizationHeader || err == ErrInvalidAuthorizationHeader {
			return echo.ErrUnauthorized
		}

		return echo.ErrInternalServerError
	}

	if err := h.usecase.verify(token); err != nil {
		if err == ErrInvalidToken || err == ErrInvalidClaims {
			return echo.ErrUnauthorized
		}

		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusOK)
}

func (h *tokenHandler) refresh(c echo.Context) error {
	token, err := getBearerTokenFromHeader(c)
	if err != nil {
		if err == ErrNoAuthorizationHeader || err == ErrInvalidAuthorizationHeader {
			return echo.ErrUnauthorized
		}

		return echo.ErrInternalServerError
	}

	newToken, err := h.usecase.refresh(token)
	if err != nil {
		if err == ErrInvalidToken || err == ErrInvalidClaims {
			return echo.ErrUnauthorized
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, generateTokenResponse{token: newToken})
}

func (h *tokenHandler) delete(c echo.Context) error {
	token, err := getBearerTokenFromHeader(c)
	if err != nil {
		if err == ErrNoAuthorizationHeader || err == ErrInvalidAuthorizationHeader {
			return echo.ErrBadRequest
		}

		return echo.ErrInternalServerError
	}

	if err := h.usecase.delete(token); err != nil {
		return echo.ErrInternalServerError
	}

	return nil
}
