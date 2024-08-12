package storage

import (
	"context"
	"errors"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage/database/dto"
	"github.com/jackc/pgx/v4"
)

func (s *BannerStorageDB) GetBannerByFeatureTag(ctx context.Context, featureID, tagID uint64) (*model.Banner, error) {
	var bannerDB dto.BannerDB
	err := s.db.Get(ctx, &bannerDB,
		`SELECT b.id, b.content, b.created_at, b.updated_at, b.is_active
				FROM banners b
				JOIN banner_tags bt ON b.id = bt.banner_id
				WHERE bt.feature_id = $1
				  AND bt.tag_id = $2`, featureID, tagID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrBannerNotFound
		}
		return nil, err
	}
	banner := dto.ConvertToBanner(bannerDB)
	return &banner, nil
}
