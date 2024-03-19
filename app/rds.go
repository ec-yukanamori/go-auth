package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var rds *redis.Client

func init() {
	rds = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RDS.Host, cfg.RDS.Port),
		Password: cfg.RDS.Password,
		DB:       cfg.RDS.DB,
	})

	ctx := context.Background()
	if err := rds.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("failed to connect to redis: %v", err))
	}
}
