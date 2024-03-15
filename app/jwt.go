package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/golang-jwt/jwt"
// )

// var secretKey = []byte(config.SecretKey)

// func generateJWT() (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"foo": "bar",
// 		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
// 	})

// 	return token.SignedString(secretKey)
// }

// func validateJWT(tokenString string) error {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return secretKey, nil
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	if _, ok := token.Claims.(jwt.MapClaims); !ok {
// 		return err
// 	}

// 	return nil
// }
