package dto

import (
	"time"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
)

type BannerVersionDTO struct {
	VersionID uint64    `json:"version_id"`
	BannerID  uint64    `json:"banner_id"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

func GetBannerVersion(banner model.BannerVersion) BannerVersionDTO {
	return BannerVersionDTO{
		VersionID: banner.VersionID,
		BannerID:  banner.BannerID,
		Content:   banner.Content,
		UpdatedAt: banner.UpdatedAt,
		IsActive:  banner.IsActive,
	}
}

func GetBannerVersionSlice(banners []model.BannerVersion) []BannerVersionDTO {
	bannerVersions := make([]BannerVersionDTO, 0, len(banners))
	for _, banner := range banners {
		bannerVersions = append(bannerVersions, GetBannerVersion(banner))
	}
	return bannerVersions
}
