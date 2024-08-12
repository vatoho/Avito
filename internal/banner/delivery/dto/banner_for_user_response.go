package dto

import (
	"time"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
)

type BannerForUser struct {
	Content string `json:"content"`
}
type BannerForAdmin struct {
	ID        uint64    `json:"id"`
	TagIDs    []uint64  `json:"tag_ids"`
	FeatureID uint64    `json:"feature_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

func GetBannerForAdmin(banner model.Banner) BannerForAdmin {
	return BannerForAdmin{
		ID:        banner.ID,
		TagIDs:    banner.TagIDs,
		FeatureID: banner.FeatureID,
		Content:   banner.Content,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
		IsActive:  banner.IsActive,
	}
}

func GetBannerForAdminSlice(banners []model.Banner) []BannerForAdmin {
	bannersForAdmin := make([]BannerForAdmin, 0, len(banners))
	for _, banner := range banners {
		bannersForAdmin = append(bannersForAdmin, GetBannerForAdmin(banner))
	}
	return bannersForAdmin
}
