package producer

import (
	"fmt"

	"github.com/IBM/sarama"
)

func newSyncProducerKafka(brokers []string) (sarama.SyncProducer, error) {
	syncProducerConfig := sarama.NewConfig()

	syncProducerConfig.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	syncProducerConfig.Producer.RequiredAcks = sarama.WaitForAll
	syncProducerConfig.Producer.CompressionLevel = sarama.CompressionLevelDefault
	syncProducerConfig.Producer.Return.Successes = true
	syncProducerConfig.Producer.Return.Errors = true
	syncProducerConfig.Producer.Compression = sarama.CompressionGZIP

	syncProducer, err := sarama.NewSyncProducer(brokers, syncProducerConfig)
	if err != nil {
		return nil, fmt.Errorf("error in creating sync producer: %w", err)
	}

	return syncProducer, nil
}

func New(brokers []string) (*SyncProducer, error) {
	syncProducer, err := newSyncProducerKafka(brokers)
	if err != nil {
		return nil, err
	}

	producer := &SyncProducer{
		brokers:      brokers,
		syncProducer: syncProducer,
	}

	return producer, nil
}
