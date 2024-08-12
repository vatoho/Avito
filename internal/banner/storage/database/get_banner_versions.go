package storage

import (
	"context"
	"errors"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage/database/dto"
	"github.com/jackc/pgx/v4"
)

func (s *BannerStorageDB) GetBannerVersions(ctx context.Context, id uint64) ([]model.BannerVersion, error) {
	var bannerVersions []dto.BannerVersionDB
	err := s.db.Select(ctx, &bannerVersions,
		`SELECT id, banner_id, content, updated_at, is_active FROM previous_banners WHERE banner_id = $1`, id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) || len(bannerVersions) == 0 {
		return nil, err
	}
	return dto.ConvertToBannerVersionSlice(bannerVersions), nil
}
