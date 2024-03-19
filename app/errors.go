package main

import "errors"

var (
	ErrInvalidToken               = errors.New("invalid token")
	ErrInvalidClaims              = errors.New("invalid claims")
	ErrNoAuthorizationHeader      = errors.New("no authorization header")
	ErrInvalidAuthorizationHeader = errors.New("invalid authorization header")
	ErrInvalidSigningMethod       = errors.New("invalid signing method")
)
