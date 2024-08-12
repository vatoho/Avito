package storage

import (
	"context"
	"errors"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage/database/dto"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func (s *BannerStorageDB) UpdateBanner(ctx context.Context, banner model.Banner) error {
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
	var oldBanner dto.BannerDB
	err = s.db.GetTx(ctx, tx, &oldBanner, `SELECT id, updated_at, content, is_active FROM banners WHERE id = $1 FOR UPDATE`, banner.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return storage.ErrBannerNotFound
		}
		return err
	}

	bannerDB, tagsDB := dto.NewBannerTagsDB(banner)
	result, err := s.db.ExecTX(ctx, tx, `UPDATE banners SET content = $1, updated_at = $2, 
                          is_active = $3 WHERE id = $4`,
		bannerDB.Content, bannerDB.UpdatedAt, bannerDB.IsActive, bannerDB.ID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return storage.ErrBannerNotFound
	}

	if !compareBannerVersions(oldBanner, bannerDB) {
		_, err = s.db.ExecTX(ctx, tx, `INSERT INTO previous_banners (banner_id, content, is_active, updated_at)
              VALUES ($1, $2, $3, $4)`,
			oldBanner.ID, oldBanner.Content, oldBanner.IsActive, oldBanner.UpdatedAt)
		if err != nil {
			return err
		}
		var count uint64
		err = s.db.QueryRowTx(ctx, tx, "SELECT COUNT(*) FROM previous_banners WHERE banner_id = $1", oldBanner.ID).Scan(&count)
		if err != nil {
			return err
		}
		if count > 3 {
			_, err = s.db.ExecTX(ctx, tx, `
			DELETE FROM previous_banners
			WHERE id = (
				SELECT id FROM previous_banners
				WHERE banner_id = $1
				ORDER BY updated_at ASC
				LIMIT 1
			)
		`, oldBanner.ID)
			if err != nil {
				return err
			}
		}

	}

	_, err = s.db.ExecTX(ctx, tx, `DELETE FROM banner_tags WHERE banner_id = $1`, bannerDB.ID)
	if err != nil {
		return err
	}

	var id uint64
	for _, tag := range tagsDB {
		err = s.db.QueryRowTx(ctx, tx,
			`INSERT INTO banner_tags (feature_id, tag_id, banner_id)
              VALUES ($1, $2, $3) RETURNING banner_id`,
			bannerDB.FeatureID, tag.ID, bannerDB.ID).Scan(&id)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				return storage.ErrDuplicateFeatureTag
			}
			return err
		}
	}

	return s.db.Commit(ctx, tx)
}

func compareBannerVersions(old dto.BannerDB, new dto.BannerDB) bool {
	if old.Content == new.Content && old.IsActive == new.IsActive {
		return true
	}
	return false
}
