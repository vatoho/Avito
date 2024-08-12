package dto

import (
	"time"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
)

type BannerDB struct {
	ID        uint64    `db:"id"`
	FeatureID uint64    `db:"feature_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	IsActive  bool      `db:"is_active"`
}

type TagDB struct {
	ID uint64 `db:"tag_id"`
}

func NewBannerTagsDB(b model.Banner) (BannerDB, []TagDB) {
	tags := make([]TagDB, 0, len(b.TagIDs))
	for _, tag := range b.TagIDs {
		tags = append(tags, TagDB{tag})
	}
	return BannerDB{
		ID:        b.ID,
		FeatureID: b.FeatureID,
		Content:   b.Content,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		IsActive:  b.IsActive,
	}, tags

}
