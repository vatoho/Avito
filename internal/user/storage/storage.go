package storage

import (
	"context"

	"github.com/ilyushkaaa/banner-service/internal/user/model"
)

type UserStorage interface {
	GetUserByToken(ctx context.Context, token string) (*model.User, error)
}
