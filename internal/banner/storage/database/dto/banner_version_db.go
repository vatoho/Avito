package dto

import (
	"time"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
)

type BannerVersionDB struct {
	VersionID uint64    `db:"id"`
	BannerID  uint64    `db:"banner_id"`
	Content   string    `db:"content"`
	UpdatedAt time.Time `db:"updated_at"`
	IsActive  bool      `db:"is_active"`
}

func NewBannerVersionDB(banner model.Banner) BannerVersionDB {
	return BannerVersionDB{
		BannerID:  banner.ID,
		Content:   banner.Content,
		UpdatedAt: banner.UpdatedAt,
		IsActive:  banner.IsActive,
	}
}
