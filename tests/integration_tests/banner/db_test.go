//go:build integration
// +build integration

package banner

import (
	"context"
	"fmt"
	"testing"

	"github.com/ilyushkaaa/banner-service/internal/banner/model"
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/database/postgres/database"
	"github.com/ilyushkaaa/banner-service/tests/fixtures"
	"github.com/ilyushkaaa/banner-service/tests/states"
	"github.com/stretchr/testify/require"
)

func setUp(t *testing.T, db database.Database, tableNames []string) {
	t.Helper()
	ctx := context.Background()

	for _, tn := range tableNames {
		err := deleteData(ctx, db, tn)
		require.NoError(t, err)
	}

}

func deleteData(ctx context.Context, db database.Database, tableName string) error {
	_, err := db.Exec(ctx, fmt.Sprintf("DELETE FROM %s", tableName))
	return err
}

func fillDataBase(t *testing.T, db database.Database) {
	t.Helper()

	fillBanners(t, db)
	fillPreviousBanners(t, db)
	fillUsers(t, db)
}

func fillBanners(t *testing.T, db database.Database) {
	t.Helper()

	insertBanner(t, db, fixtures.Banner().Valid1().Val())
	insertBanner(t, db, fixtures.Banner().Valid2().Val())
}

func fillPreviousBanners(t *testing.T, db database.Database) {
	t.Helper()

	insertPreviousBanner(t, db, fixtures.Banner().Valid1().Content(states.Content3).Val(), uint64(1))
	insertPreviousBanner(t, db, fixtures.Banner().Valid1().Content(states.Content4).Val(), uint64(2))
}

func insertBanner(t *testing.T, db database.Database, banner model.Banner) {
	t.Helper()
	ctx := context.Background()
	_, err := db.Exec(ctx,
		`INSERT INTO banners (id, content, created_at, updated_at, is_active)
              VALUES ($1, $2, $3, $4, $5)`,
		banner.ID, banner.Content, banner.CreatedAt, banner.UpdatedAt, banner.IsActive)
	require.NoError(t, err)

	for _, tag := range banner.TagIDs {
		_, err = db.Exec(ctx,
			`INSERT INTO banner_tags (feature_id, tag_id, banner_id)
              VALUES ($1, $2, $3) RETURNING banner_id`,
			banner.FeatureID, tag, banner.ID)
		require.NoError(t, err)
	}
}

func fillUsers(t *testing.T, db database.Database) {
	t.Helper()
	ctx := context.Background()
	_, err := db.Exec(ctx,
		`INSERT INTO users (tag_id, token, role)
              VALUES ($1, $2, $3)`,
		states.TagID1, states.TokenUser, states.RoleUser)
	require.NoError(t, err)

	_, err = db.Exec(ctx,
		`INSERT INTO users (tag_id, token, role)
              VALUES ($1, $2, $3)`,
		states.TagID2, states.TokenAdmin, states.RoleAdmin)
	require.NoError(t, err)

}

func insertPreviousBanner(t *testing.T, db database.Database, banner model.Banner, id uint64) {
	t.Helper()
	ctx := context.Background()
	_, err := db.Exec(ctx,
		`INSERT INTO previous_banners (id, content, banner_id, updated_at, is_active)
              VALUES ($1, $2, $3, $4, $5)`,
		id, banner.Content, banner.ID, banner.UpdatedAt, banner.IsActive)
	require.NoError(t, err)
}
