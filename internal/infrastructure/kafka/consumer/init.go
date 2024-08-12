package consumer

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/kafka"
)

func newConsumerGroup(cfg *kafka.ConfigKafka) (sarama.ConsumerGroup, error) {
	config := cfg.ConsumerConfig
	if config == nil {
		config = defaultConsumerConfig()
	}
	consumer, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.ConsumerGroupID, config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func defaultConsumerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Group.ResetInvalidOffsets = true
	config.Consumer.Group.Heartbeat.Interval = 3 * time.Second
	config.Consumer.Group.Session.Timeout = 60 * time.Second
	config.Consumer.Group.Rebalance.Timeout = 60 * time.Second
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}
	return config
}
