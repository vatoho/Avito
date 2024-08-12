//go:build integration
// +build integration

package banner

import (
	"context"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/ilyushkaaa/banner-service/internal/banner/delivery"
	serviceBanner "github.com/ilyushkaaa/banner-service/internal/banner/service"
	storageBanner "github.com/ilyushkaaa/banner-service/internal/banner/storage/database"
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/database/postgres/database"
	redisApp "github.com/ilyushkaaa/banner-service/internal/infrastructure/database/redis"
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/kafka"
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/kafka/consumer"
	"github.com/ilyushkaaa/banner-service/internal/middleware"
	serviceUser "github.com/ilyushkaaa/banner-service/internal/user/service"
	storageUser "github.com/ilyushkaaa/banner-service/internal/user/storage/database"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type BannerTestFixtures struct {
	del       *delivery.BannerDelivery
	mw        *middleware.Middleware
	db        *database.PGDatabase
	redisConn redis.Conn
	cfg       *kafka.ConfigKafka
	cancel    context.CancelFunc
}

func newBannerTestFixtures(t *testing.T) BannerTestFixtures {
	t.Helper()

	logger := zap.NewNop().Sugar()

	ctx, cancel := context.WithCancel(context.Background())
	db, err := database.New(ctx)
	require.NoError(t, err)

	redisConn, err := redisApp.Init()
	require.NoError(t, err)

	cfg, err := kafka.NewConfig()
	require.NoError(t, err)

	stBanner := storageBanner.New(db, redisConn, logger)
	svBanner := serviceBanner.New(stBanner, cfg.Producer)
	d := delivery.New(svBanner, logger)

	stUser := storageUser.New(db)
	svUser := serviceUser.New(stUser)

	mw := middleware.New(svUser, logger)

	go func() {
		err = consumer.Run(ctx, cfg, stBanner, logger)
		if err != nil {
			logger.Errorf("error in consumer running")
		}
	}()
	return BannerTestFixtures{
		del:       d,
		mw:        mw,
		db:        db,
		redisConn: redisConn,
		cfg:       cfg,
		cancel:    cancel,
	}
}

func (b *BannerTestFixtures) Close(t *testing.T) {
	b.cancel()
	err := b.db.Close()
	require.NoError(t, err)
	err = b.cfg.Close()
	require.NoError(t, err)
	err = b.redisConn.Close()

}
