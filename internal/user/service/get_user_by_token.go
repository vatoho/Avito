package service

import (
	"context"

	"github.com/ilyushkaaa/banner-service/internal/pkg/hash"
	"github.com/ilyushkaaa/banner-service/internal/user/model"
)

func (s *UserServiceApp) GetUserByToken(ctx context.Context, token string) (*model.User, error) {
	tokenHashed, err := hash.GetHash(token)
	if err != nil {
		return nil, err
	}
	return s.storage.GetUserByToken(ctx, tokenHashed)
}
