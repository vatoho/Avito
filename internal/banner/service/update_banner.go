package service

import (
	"context"
	"time"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
)

type BannerToUpdate struct {
	ID        uint64
	TagIDs    *[]uint64
	FeatureID *uint64
	Content   *string
	IsActive  *bool
}

func (s *BannerServiceApp) UpdateBanner(ctx context.Context, banner BannerToUpdate) error {
	currentBannerVersion, err := s.storage.GetBannerByID(ctx, banner.ID)
	if err != nil {
		return err
	}
	feature, tags, err := s.storage.GetBannerFeatureTags(ctx, banner.ID)
	if err != nil {
		return err
	}

	currentBannerVersion.FeatureID = feature
	currentBannerVersion.TagIDs = tags

	bannerToUpdate := constructBannerToUpdate(banner, *currentBannerVersion)
	bannerToUpdate.UpdatedAt = time.Now()
	return s.storage.UpdateBanner(ctx, bannerToUpdate)
}

func constructBannerToUpdate(newBanner BannerToUpdate, oldBanner model.Banner) model.Banner {
	if newBanner.TagIDs != nil {
		oldBanner.TagIDs = *newBanner.TagIDs
	}
	if newBanner.FeatureID != nil {
		oldBanner.FeatureID = *newBanner.FeatureID
	}
	if newBanner.Content != nil {
		oldBanner.Content = *newBanner.Content
	}
	if newBanner.IsActive != nil {
		oldBanner.IsActive = *newBanner.IsActive
	}
	return oldBanner
}
