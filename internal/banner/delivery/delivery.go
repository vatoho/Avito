package delivery

import (
	"github.com/ilyushkaaa/banner-service/internal/banner/service"
	"go.uber.org/zap"
)

type BannerDelivery struct {
	service service.BannerService
	logger  *zap.SugaredLogger
}

func New(service service.BannerService, logger *zap.SugaredLogger) *BannerDelivery {
	return &BannerDelivery{
		service: service,
		logger:  logger,
	}
}
