package storage

import (
	"context"
	"errors"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage/database/dto"
	"github.com/jackc/pgconn"
)

func (s *BannerStorageDB) AddBanner(ctx context.Context, banner model.Banner) (*model.Banner, error) {
	bannerDB, tagsDB := dto.NewBannerTagsDB(banner)
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
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

	var id uint64
	err = s.db.QueryRowTx(ctx, tx,
		`INSERT INTO banners (content, created_at, updated_at, is_active)
              VALUES ($1, $2, $3, $4) RETURNING id`,
		bannerDB.Content, bannerDB.CreatedAt, bannerDB.UpdatedAt, bannerDB.IsActive).Scan(&id)
	if err != nil {
		return nil, err
	}

	bannerDB.ID = id
	banner.ID = id
	for _, tag := range tagsDB {
		err = s.db.QueryRowTx(ctx, tx,
			`INSERT INTO banner_tags (feature_id, tag_id, banner_id)
              VALUES ($1, $2, $3) RETURNING banner_id`,
			bannerDB.FeatureID, tag.ID, id).Scan(&id)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				return nil, storage.ErrDuplicateFeatureTag
			}
			return nil, err
		}
	}

	err = s.db.Commit(ctx, tx)
	if err != nil {
		return nil, err
	}

	return &banner, nil
}
