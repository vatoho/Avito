package storage

import (
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
	"golang.org/x/net/context"
)

func (s *BannerStorageDB) DeleteBanner(ctx context.Context, id uint64) error {
	result, err := s.db.Exec(ctx, `DELETE FROM banners WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return storage.ErrBannerNotFound
	}
	return nil
}
