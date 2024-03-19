package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var secret []byte

func init() {
	secret = make([]byte, 32)
	_, err := rand.Read(secret)
	if err != nil {
		panic(fmt.Errorf("failed to generate random secret: %v", err))
	}
}

type customClaims struct {
	jwt.StandardClaims
	Roles []string `json:"roles,omitempty"`
}

func createToken(input generateTokenInput) (string, string, error) {
	claims := customClaims{
		StandardClaims: jwt.StandardClaims{
			Audience: input.AppURI,
			// TODO: Make it possible to configure for each system
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(cfg.JWT.ExpirationMinutes)).Unix(),
			Id:        uuid.NewString(),
			IssuedAt:  time.Now().Unix(),
			Subject:   input.UserID,
		},
		Roles: input.Roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	return claims.Id, tokenString, nil
}

func verifyToken(tokenString string) (customClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}

		return secret, nil
	})
	if err != nil {
		if _, ok := err.(*jwt.ValidationError); ok {
			return customClaims{}, ErrInvalidToken
		}

		return customClaims{}, err
	}

	if !token.Valid {
		return customClaims{}, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return customClaims{}, ErrInvalidClaims
	}

	return mapClaims2CustomClaims(claims), nil
}

func mapClaims2CustomClaims(claims jwt.MapClaims) customClaims {
	var custom customClaims

	if val, ok := claims["exp"].(float64); ok {
		custom.ExpiresAt = int64(val)
	}

	if val, ok := claims["iat"].(float64); ok {
		custom.IssuedAt = int64(val)
	}

	if val, ok := claims["nbf"].(float64); ok {
		custom.NotBefore = int64(val)
	}

	if val, ok := claims["iss"].(string); ok {
		custom.Issuer = val
	}

	if val, ok := claims["sub"].(string); ok {
		custom.Subject = val
	}

	if val, ok := claims["aud"].(string); ok {
		custom.Audience = val
	}

	if val, ok := claims["jti"].(string); ok {
		custom.Id = val
	}

	if val, ok := claims["roles"].([]interface{}); ok {
		for _, role := range val {
			if str, ok := role.(string); ok {
				custom.Roles = append(custom.Roles, str)
			}
		}
	}

	return custom
}

func getJWTID(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, nil)
	if err != nil {
		return "", err
	}

	if token == nil {
		return "", errors.New("token is nil")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("cannot get claims from token")
	}

	jti, ok := claims["jti"].(string)
	if !ok {
		return "", errors.New("jti claim not found")
	}

	return jti, nil
}
