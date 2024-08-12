package storage

import (
	"context"
	"errors"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage/database/dto"
	"github.com/jackc/pgx/v4"
)

func (s *BannerStorageDB) GetBannerByID(ctx context.Context, id uint64) (*model.Banner, error) {
	var bannerDB dto.BannerDB
	err := s.db.Get(ctx, &bannerDB,
		`SELECT id, content, created_at, updated_at, is_active 
                FROM banners WHERE id = $1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrBannerNotFound
		}
		return nil, err
	}
	banner := dto.ConvertToBanner(bannerDB)
	return &banner, nil
}
