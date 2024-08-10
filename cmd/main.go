package main

import (
	"api_gateway/api"
	"api_gateway/configs"
	"api_gateway/grpc/client"
	"api_gateway/pkg/logger"
	"api_gateway/storage"
	"api_gateway/storage/redis"
	"context"

	"go.uber.org/zap"
)

func main() {
	config := configs.Load()

	logger := logger.NewLogger(config.ServiceName, config.LoggerLevel, config.LogPath)

	services, err := client.NewGrpcClients(config)
	if err != nil {
		logger.Fatal("Failed make client connections ", zap.Error(err))
		return
	}

	redisClient, err := redis.NewRedisClient()
	if err != nil {
		logger.Fatal("Failed while creating redis client ", zap.Error(err))
		return
	}

	storage, err := storage.New(context.Background(), config, &logger)

	router := api.NewRouter(logger, services, storage)

	logger.Info("Gin router is running..")
	err = router.Run(config.ApiGatewayHttpHost + config.ApiGatewayHttpPort)
	if err != nil {
		logger.Fatal("Gin router failed to run", zap.Error(err))
		return
	}
}
