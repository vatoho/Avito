package storage

import (
	"context"
	"fmt"

	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
)

func (s *BannerStorageDB) DeleteBannersByFeatureTag(ctx context.Context, featureID, tagID uint64) error {
	query := `DELETE FROM banners
				USING banner_tags
				WHERE banner_tags.banner_id = banners.id`

	if featureID != 0 {
		query += fmt.Sprintf(" AND banner_tags.feature_id = %d", featureID)
	}
	if tagID != 0 {
		query += fmt.Sprintf(" AND banner_tags.tag_id = %d", tagID)
	}
	result, err := s.db.Exec(ctx, query)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return storage.ErrBannerNotFound
	}
	return nil
}
