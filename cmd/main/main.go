package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ilyushkaaa/banner-service/internal/banner/delivery"
	serviceBanner "github.com/ilyushkaaa/banner-service/internal/banner/service"
	storageBanner "github.com/ilyushkaaa/banner-service/internal/banner/storage/database"
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/database/postgres/database"
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/database/redis"
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/kafka"
	"github.com/ilyushkaaa/banner-service/internal/infrastructure/kafka/consumer"
	"github.com/ilyushkaaa/banner-service/internal/middleware"
	"github.com/ilyushkaaa/banner-service/internal/routes"
	serviceUser "github.com/ilyushkaaa/banner-service/internal/user/service"
	storageUser "github.com/ilyushkaaa/banner-service/internal/user/storage/database"
	"go.uber.org/zap"
)

func main() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("error in logger initialization: %v", err)
	}
	logger := zapLogger.Sugar()
	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Printf("error in logger sync: %v", err)
		}
	}()
	//err = godotenv.Load(".env.testing")
	//if err != nil {
	//	logger.Fatalf("error in getting env: %s", err)
	//}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := database.New(ctx)
	if err != nil {
		logger.Fatalf("error in database init: %v", err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			logger.Errorf("error in closing db")
		}
	}()

	redisConn, err := redis.Init()
	if err != nil {
		logger.Fatalf("error on connection to redis: %v", err)
	}
	defer func() {
		err = redisConn.Close()
		if err != nil {
			logger.Infof("error on redis close: %s", err.Error())
		}
	}()

	cfg, err := kafka.NewConfig()
	if err != nil {
		logger.Fatalf("error in kafka config init: %v", err)
	}
	defer func() {
		err = cfg.Close()
		if err != nil {
			logger.Errorf("error in closing sync kafka producer: %v", err)
		}
	}()

	stBanner := storageBanner.New(db, redisConn, logger)
	svBanner := serviceBanner.New(stBanner, cfg.Producer)
	d := delivery.New(svBanner, logger)

	stUser := storageUser.New(db)
	svUser := serviceUser.New(stUser)

	mw := middleware.New(svUser, logger)
	router := routes.GetRouter(d, mw)

	go func() {
		err = consumer.Run(ctx, cfg, stBanner, logger)
		if err != nil {
			logger.Errorf("error in consumer running")
		}
	}()

	port := os.Getenv("APP_PORT")
	addr := ":" + port
	logger.Infow("starting server",
		"type", "START",
		"addr", addr,
	)
	logger.Fatal(http.ListenAndServe(addr, router))
}
