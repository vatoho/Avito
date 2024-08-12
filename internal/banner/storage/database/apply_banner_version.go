package storage

import (
	"context"
	"errors"

	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage/database/dto"
	"github.com/jackc/pgx/v4"
)

func (s *BannerStorageDB) ApplyBannerVersion(ctx context.Context, versionID uint64) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			err = s.db.Rollback(ctx, tx)
			if err != nil {
				s.logger.Errorf("error in transaction rollback: %v", err)
				return
			}
		}
	}()

	var bannerVersionDB dto.BannerVersionDB
	err = s.db.GetTx(ctx, tx, &bannerVersionDB,
		`SELECT banner_id, content, updated_at, is_active 
                FROM previous_banners WHERE id = $1 FOR UPDATE`, versionID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return storage.ErrNoBannerVersions
		}
		return err
	}
	var bannerDB dto.BannerDB
	err = s.db.GetTx(ctx, tx, &bannerDB,
		`SELECT id, content, updated_at, is_active 
                FROM banners WHERE id = $1 FOR UPDATE`, bannerVersionDB.BannerID)
	if err != nil {
		return err
	}

	_, err = s.db.ExecTX(ctx, tx, `UPDATE banners SET content = $1, updated_at = $2, 
                          is_active = $3 WHERE id = $4`,
		bannerVersionDB.Content, bannerVersionDB.UpdatedAt, bannerVersionDB.IsActive, bannerVersionDB.BannerID)
	if err != nil {
		return err
	}
	_, err = s.db.ExecTX(ctx, tx, `UPDATE previous_banners SET content = $1, updated_at = $2, 
                          is_active = $3 WHERE id = $4`,
		bannerDB.Content, bannerDB.UpdatedAt, bannerDB.IsActive, versionID)
	if err != nil {
		return err
	}
	return s.db.Commit(ctx, tx)

}
