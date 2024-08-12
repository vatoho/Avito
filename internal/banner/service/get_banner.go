package service

import (
	"context"

	"github.com/ilyushkaaa/banner-service/internal/banner/filter"
	"github.com/ilyushkaaa/banner-service/internal/banner/model"
)

func (s *BannerServiceApp) GetBanners(ctx context.Context, filter filter.Filter) ([]model.Banner, error) {
	banners, err := s.storage.GetBanners(ctx, filter)
	if err != nil {
		return nil, err
	}
	for i, banner := range banners {
		feature, tags, err := s.storage.GetBannerFeatureTags(ctx, banner.ID)
		if err != nil {
			return nil, err
		}
		banners[i].FeatureID = feature
		banners[i].TagIDs = tags
	}
	return banners, nil
}
