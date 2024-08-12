package workers

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/ilyushkaaa/banner-service/internal/banner/service/workers/dto"
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/kafka/producer"
)

type DeleteBannersProducer struct {
	producer producer.Producer
	topic    string
}

type SendMessageResult struct {
	Partition int32
	Offset    int64
	Error     error
}

func NewProducer(producer producer.Producer, topic string) *DeleteBannersProducer {
	return &DeleteBannersProducer{
		producer: producer,
		topic:    topic,
	}
}

func (s *DeleteBannersProducer) SendMessage(featureTag dto.FeatureTagDTO) SendMessageResult {
	kafkaMsg, err := s.BuildMessage(featureTag)
	if err != nil {
		return SendMessageResult{Error: err}
	}

	partition, offset, err := s.producer.SendMessage(kafkaMsg)

	if err != nil {
		return SendMessageResult{Error: err}
	}
	return SendMessageResult{
		Partition: partition,
		Offset:    offset,
		Error:     err}
}

func (s *DeleteBannersProducer) BuildMessage(featureTag dto.FeatureTagDTO) (*sarama.ProducerMessage, error) {
	msg, err := json.Marshal(featureTag)

	if err != nil {
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(msg),
		Partition: -1,
	}, nil
}

func (s *DeleteBannersProducer) Close() error {
	return s.producer.Close()
}
