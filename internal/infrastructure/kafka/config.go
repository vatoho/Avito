package kafka

import (
	"os"
	"strings"

	"github.com/IBM/sarama"
	"github.com/ilyushkaaa/banner-service/internal/banner/service/workers"
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/kafka/producer"
)

type ConfigKafka struct {
	Brokers         []string
	ConsumerGroupID string
	Topic           string
	Producer        *workers.DeleteBannersProducer
	ConsumerConfig  *sarama.Config
}

func NewConfig() (*ConfigKafka, error) {
	brokersFromEnv := os.Getenv("KAFKA_BROKERS")
	brokers := strings.Split(brokersFromEnv, ",")
	syncProducer, err := producer.New(brokers)
	if err != nil {
		return nil, err
	}
	topic := os.Getenv("KAFKA_EVENTS_TOPIC")
	groupID := os.Getenv("EVENTS_CONSUMER_GROUP_ID")
	delProd := workers.NewProducer(syncProducer, topic)
	return &ConfigKafka{
		Brokers:         brokers,
		ConsumerGroupID: groupID,
		Topic:           topic,
		Producer:        delProd,
	}, nil
}

func (c *ConfigKafka) Close() error {
	return c.Producer.Close()
}
