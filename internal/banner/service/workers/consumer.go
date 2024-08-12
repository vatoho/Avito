package workers

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/ilyushkaaa/banner-service/internal/banner/service/workers/dto"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
	"go.uber.org/zap"
)

type DeleteWorker struct {
	storage storage.BannerStorage
	logger  *zap.SugaredLogger
	ready   chan bool
}

func NewDeleteWorker(logger *zap.SugaredLogger, storage storage.BannerStorage) DeleteWorker {
	return DeleteWorker{
		logger:  logger,
		storage: storage,
		ready:   make(chan bool),
	}
}

func (s *DeleteWorker) Ready() <-chan bool {
	return s.ready
}

func (s *DeleteWorker) Setup(_ sarama.ConsumerGroupSession) error {
	close(s.ready)

	return nil
}

func (s *DeleteWorker) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (s *DeleteWorker) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():

			var featureTag dto.FeatureTagDTO
			err := json.Unmarshal(message.Value, &featureTag)
			if err != nil {
				s.logger.Errorf("can not unmarshal message: %v", err)
				continue
			}

			err = s.storage.DeleteBannersByFeatureTag(session.Context(), featureTag.FeatureID, featureTag.TagID)
			if err != nil {
				s.logger.Errorf("error in deleting banners: %v", err)
			}

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
