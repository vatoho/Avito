package middleware

import (
	"github.com/ilyushkaaa/banner-service/internal/user/service"
	"go.uber.org/zap"
)

type Middleware struct {
	userService service.UserService
	logger      *zap.SugaredLogger
}

func New(userService service.UserService, logger *zap.SugaredLogger) *Middleware {
	return &Middleware{
		userService: userService,
		logger:      logger,
	}
}
