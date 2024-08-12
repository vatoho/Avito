package service

import (
	"github.com/ilyushkaaa/banner-service/internal/banner/service/workers"
	"github.com/ilyushkaaa/banner-service/internal/banner/service/workers/dto"
)

func (s *BannerServiceApp) DeleteBannersByFeatureTag(featureID, tagID uint64) workers.SendMessageResult {
	featureTag := dto.FeatureTagDTO{FeatureID: featureID, TagID: tagID}
	return s.deleteProducer.SendMessage(featureTag)
}
