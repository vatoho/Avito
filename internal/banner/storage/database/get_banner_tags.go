package storage

import (
	"context"
	"errors"

	"github.com/ilyushkaaa/banner-service/internal/banner/storage/database/dto"
	"github.com/jackc/pgx/v4"
)

func (s *BannerStorageDB) GetBannerFeatureTags(ctx context.Context, id uint64) (uint64, []uint64, error) {
	var featureTags []dto.FeatureTag
	err := s.db.Select(ctx, &featureTags,
		`SELECT feature_id, tag_id FROM banner_tags WHERE banner_id = $1`, id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) || len(featureTags) == 0 {
		return 0, nil, err
	}
	tags := make([]uint64, 0, len(featureTags))
	for _, ft := range featureTags {
		tags = append(tags, ft.TagID)
	}
	return featureTags[0].FeatureID, tags, nil
}
