package service

import (
	"context"

	"github.com/ilyushkaaa/banner-service/internal/user/model"
	"github.com/ilyushkaaa/banner-service/internal/user/storage"
)

type UserService interface {
	GetUserByToken(ctx context.Context, token string) (*model.User, error)
}

type UserServiceApp struct {
	storage storage.UserStorage
}

func New(storage storage.UserStorage) *UserServiceApp {
	return &UserServiceApp{storage: storage}
}
