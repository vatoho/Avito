package storage

import (
	"context"

	"github.com/ilyushkaaa/banner-service/internal/banner/filter"
	"github.com/ilyushkaaa/banner-service/internal/banner/model"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage/database/dto"
)

type BannerStorage interface {
	AddBanner(ctx context.Context, banner model.Banner) (*model.Banner, error)
	GetBanners(ctx context.Context, filter filter.Filter) ([]model.Banner, error)
	GetBannerByID(ctx context.Context, ID uint64) (*model.Banner, error)
	GetBannerByFeatureTag(ctx context.Context, featureID, tagID uint64) (*model.Banner, error)
	GetBannerFromCache(featureID, tagID uint64) (*dto.BannerFromCache, error)
	SaveBannerToCache(bannerFromCache dto.BannerFromCache, featureID, tagID uint64)
	GetBannerFeatureTags(ctx context.Context, ID uint64) (uint64, []uint64, error)
	UpdateBanner(ctx context.Context, banner model.Banner) error
	DeleteBanner(ctx context.Context, ID uint64) error
	GetBannerVersions(ctx context.Context, ID uint64) ([]model.BannerVersion, error)
	ApplyBannerVersion(ctx context.Context, versionID uint64) error
	DeleteBannersByFeatureTag(ctx context.Context, featureID, tagID uint64) error
}
