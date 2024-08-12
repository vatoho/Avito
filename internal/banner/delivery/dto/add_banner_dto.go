package dto

import (
	"github.com/asaskevich/govalidator"
	"github.com/ilyushkaaa/banner-service/internal/banner/model"
)

type BannerToAdd struct {
	TagIDs    []uint64 `json:"tag_ids" valid:"required"`
	FeatureID uint64   `json:"feature_id" valid:"required"`
	Content   string   `json:"content" valid:"required,json"`
	IsActive  bool     `json:"is_active"`
}

func (b *BannerToAdd) Validate() error {
	_, err := govalidator.ValidateStruct(b)
	return err
}

func ConvertToBanner(b BannerToAdd) model.Banner {
	return model.Banner{
		TagIDs:    b.TagIDs,
		FeatureID: b.FeatureID,
		Content:   b.Content,
		IsActive:  b.IsActive,
	}
}
