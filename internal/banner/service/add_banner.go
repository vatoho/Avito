package service

import (
	"context"
	"time"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
)

func (s *BannerServiceApp) AddBanner(ctx context.Context, banner model.Banner) (*model.Banner, error) {
	creationTime := time.Now()
	banner.CreatedAt = creationTime
	banner.UpdatedAt = creationTime

	addedBanner, err := s.storage.AddBanner(ctx, banner)
	if err != nil {
		return nil, err
	}
	return addedBanner, nil
}
