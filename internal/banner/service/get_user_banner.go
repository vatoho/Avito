package service

import (
	"context"

	"github.com/ilyushkaaa/banner-service/internal/banner/storage/database/dto"
)

func (s *BannerServiceApp) GetUserBanner(ctx context.Context, featureID, tagID uint64, lastVersion bool) (string, error) {
	if !lastVersion {
		bannerCache, err := s.storage.GetBannerFromCache(featureID, tagID)
		if err == nil && bannerCache != nil {
			return bannerCache.Content, bannerCache.Error
		}
	}

	banner, err := s.storage.GetBannerByFeatureTag(ctx, featureID, tagID)
	if err != nil {
		return "", err
	}

	if !banner.IsActive {
		return "", ErrBannerIsInactive
	}

	if !lastVersion {
		go s.storage.SaveBannerToCache(dto.BannerFromCache{Content: banner.Content, Error: err}, featureID, tagID)
	}

	return banner.Content, nil
}
