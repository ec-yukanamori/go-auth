package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// func main() {
// 	e := echo.New()

// 	e.Use(middleware.Recover())
// 	e.Use(middleware.RequestID())
// 	e.Use(middleware.RequestLoggerWithConfig(requestLoggerConfig))

// 	// TODO: extract routing
// 	e.GET("/ping", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "pong")
// 	})

// 	if err := e.Start(fmt.Sprintf(":%s", config.Server.Port)); err != nil {
// 		panic(fmt.Sprintf("failed to start server: %v", err))
// 	}
// }

func main() {
	ExampleClient()
}

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
