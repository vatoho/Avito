package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/ilyushkaaa/banner-service/internal/banner/filter"
	"github.com/ilyushkaaa/banner-service/internal/banner/model"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage/database/dto"
	"github.com/jackc/pgx/v4"
)

func (s *BannerStorageDB) GetBanners(ctx context.Context, filter filter.Filter) ([]model.Banner, error) {
	var bannersDB []dto.BannerDB
	query := `SELECT DISTINCT b.id, b.content, b.created_at, b.updated_at, b.is_active 
			   FROM banners b
			   JOIN banner_tags bt ON b.id = bt.banner_id
			   WHERE 1=1`

	if filter.FeatureID != 0 {
		query += fmt.Sprintf(" AND bt.feature_id = %d", filter.FeatureID)
	}
	if filter.TagID != 0 {
		query += fmt.Sprintf(" AND bt.tag_id = %d", filter.TagID)
	}
	if filter.Offset != 0 {
		query += fmt.Sprintf(" OFFSET %d", filter.Offset)
	}
	if filter.Limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", filter.Limit)
	}

	err := s.db.Select(ctx, &bannersDB, query)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	return dto.ConvertToBannerSlice(bannersDB), nil
}
