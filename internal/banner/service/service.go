package service

import (
	"context"

	"github.com/ilyushkaaa/banner-service/internal/banner/filter"
	"github.com/ilyushkaaa/banner-service/internal/banner/model"
	"github.com/ilyushkaaa/banner-service/internal/banner/service/workers"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
)

type BannerService interface {
	AddBanner(ctx context.Context, banner model.Banner) (*model.Banner, error)
	GetBanners(ctx context.Context, filter filter.Filter) ([]model.Banner, error)
	GetUserBanner(ctx context.Context, featureID, tagID uint64, lastVersion bool) (string, error)
	UpdateBanner(ctx context.Context, banner BannerToUpdate) error
	DeleteBanner(ctx context.Context, ID uint64) error
	GetBannerVersions(ctx context.Context, ID uint64) ([]model.BannerVersion, error)
	ApplyBannerVersion(ctx context.Context, versionID uint64) error
	DeleteBannersByFeatureTag(featureID, tagID uint64) workers.SendMessageResult
}

type BannerServiceApp struct {
	storage        storage.BannerStorage
	deleteProducer *workers.DeleteBannersProducer
}

func New(storage storage.BannerStorage, deleteProducer *workers.DeleteBannersProducer) *BannerServiceApp {
	return &BannerServiceApp{
		storage:        storage,
		deleteProducer: deleteProducer,
	}
}
