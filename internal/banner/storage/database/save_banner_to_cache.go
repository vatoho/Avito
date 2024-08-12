package storage

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage/database/dto"
)

func (s *BannerStorageDB) SaveBannerToCache(bannerCache dto.BannerFromCache, featureID, tagID uint64) {
	bannerCacheJSON, err := json.Marshal(bannerCache)
	key := constructRedisKey(featureID, tagID)
	result, err := redis.String(s.redisConn.Do("SET", key, bannerCacheJSON, "EX", s.cacheExpireTime))
	if err != nil || result != "OK" {
		s.logger.Errorf("error in saving banner into cache: %s", err)
	}
}
