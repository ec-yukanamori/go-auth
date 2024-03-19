package main

import (
	"go.uber.org/zap"
)

type (
	tokenUsecase interface {
		generate(input generateTokenInput) (string, error)
		verify(token string) error
		refresh(token string) (string, error)
		delete(token string) error
	}

	tokenUsecaseImpl struct {
		repo tokenRepository
	}
)

func newtokenUsecase(tokenRepo tokenRepository) *tokenUsecaseImpl {
	return &tokenUsecaseImpl{
		repo: tokenRepo,
	}
}

func (u *tokenUsecaseImpl) generate(input generateTokenInput) (string, error) {
	id, token, err := createToken(input)
	if err != nil {
		logger.Error("failed to create token", zap.Error(err))
		return "", err
	}

	if err := u.repo.store(id, token); err != nil {
		return "", err
	}

	return token, nil
}

func (u *tokenUsecaseImpl) verify(token string) error {
	claims, err := verifyToken(token)
	if err != nil {
		if err == ErrInvalidToken || err == ErrInvalidClaims {
			return err
		}

		logger.Error("failed to verify token", zap.Error(err))
		return err
	}

	return u.repo.verify(claims.Id)
}

func (u *tokenUsecaseImpl) refresh(token string) (string, error) {
	claims, err := verifyToken(token)
	if err != nil {
		if err == ErrInvalidToken || err == ErrInvalidClaims {
			return "", err
		}

		logger.Error("failed to verify token", zap.Error(err))
		return "", err
	}

	if err := u.repo.verify(claims.Id); err != nil {
		return "", err
	}

	input := generateTokenInput{
		AppURI: claims.Audience,
		UserID: claims.Subject,
		Roles:  claims.Roles,
	}

	// FIXME: This is a bug. We should not delete the old token before storing the new one.
	newID, newToken, err := createToken(input)
	if err != nil {
		logger.Error("failed to refresh token", zap.Error(err))
		return "", err
	}

	if err := u.repo.store(newID, newToken); err != nil {
		return "", err
	}

	if err := u.repo.delete(claims.Id); err != nil {
		return "", err
	}

	return newToken, nil
}

func (u *tokenUsecaseImpl) delete(token string) error {
	id, err := getJWTID(token)
	if err != nil {
		logger.Error("failed to get jwt id", zap.Error(err))
		return err
	}

	return u.repo.delete(id)
}
