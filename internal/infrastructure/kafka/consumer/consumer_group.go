package consumer

import (
	"context"
	"sync"

	"github.com/ilyushkaaa/banner-service/internal/banner/service/workers"
	"github.com/ilyushkaaa/banner-service/internal/banner/storage"
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/kafka"
	"go.uber.org/zap"
)

func Run(ctx context.Context, cfg *kafka.ConfigKafka, st storage.BannerStorage, logger *zap.SugaredLogger) error {
	client, err := newConsumerGroup(cfg)
	if err != nil {
		return err
	}

	deleteWorker := workers.NewDeleteWorker(logger, st)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err = client.Consume(ctx, []string{cfg.Topic}, &deleteWorker); err != nil {
				logger.Errorf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-deleteWorker.Ready()
	logger.Info("Sarama consumer up and running!...")

	wg.Wait()

	if err = client.Close(); err != nil {
		return err
	}
	return nil
}
