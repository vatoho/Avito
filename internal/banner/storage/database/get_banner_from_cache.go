package storage

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage/database/dto"
)

func (s *BannerStorageDB) GetBannerFromCache(featureID, tagID uint64) (*dto.BannerFromCache, error) {
	var bannerCash dto.BannerFromCache
	key := constructRedisKey(featureID, tagID)
	bannerFromRedis, err := redis.String(s.redisConn.Do("GET", key))
	if err != nil {
		s.logger.Errorf("redis error: %v", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(bannerFromRedis), &bannerCash)
	if err != nil {
		s.logger.Errorf("unmarshall error: %v", err)
		return nil, err
	}
	return &bannerCash, nil
}

func constructRedisKey(featureID, tagID uint64) string {
	return strings.Join([]string{strconv.FormatUint(featureID, 10), strconv.FormatUint(tagID, 10)}, "_")
}
