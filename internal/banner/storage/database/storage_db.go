package storage

import (
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/database/postgres/database"
	"go.uber.org/zap"

	"github.com/gomodule/redigo/redis"
)

const expireTime = 5 * 60

type BannerStorageDB struct {
	db              database.Database
	redisConn       redis.Conn
	cacheExpireTime int
	logger          *zap.SugaredLogger
}

func New(db database.Database, redisConn redis.Conn, logger *zap.SugaredLogger) *BannerStorageDB {
	return &BannerStorageDB{
		db:              db,
		logger:          logger,
		redisConn:       redisConn,
		cacheExpireTime: expireTime,
	}
}
