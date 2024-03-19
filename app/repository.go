package main

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type (
	tokenRepository interface {
		store(id, token string) error
		verify(id string) error
		delete(id string) error
	}

	tokenRepositoryImpl struct {
		rds *redis.Client
		ctx context.Context
	}
)

func newTokenRepository(rds *redis.Client) *tokenRepositoryImpl {
	return &tokenRepositoryImpl{
		rds: rds,
		ctx: context.Background(),
	}
}

func (r *tokenRepositoryImpl) store(id, token string) error {
	if err := r.rds.Set(r.ctx, id, token, 0).Err(); err != nil {
		logger.Error("failed to store token", zap.Error(err))
		return err
	}

	return nil
}

func (r *tokenRepositoryImpl) verify(id string) error {
	_, err := r.rds.Get(r.ctx, id).Result()
	if err != nil {
		return ErrInvalidToken
	}

	return nil
}

func (r *tokenRepositoryImpl) delete(id string) error {
	if err := r.rds.Del(r.ctx, id).Err(); err != nil {
		logger.Error("failed to delete token", zap.Error(err))
		return err
	}

	return nil
}
