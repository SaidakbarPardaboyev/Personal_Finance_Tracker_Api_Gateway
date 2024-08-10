package main

import (
	"api_gateway/api"
	"api_gateway/configs"
	"api_gateway/grpc/client"
	"api_gateway/pkg/logger"

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

	router := api.NewRouter(logger, services)

	logger.Info("Gin router is running..")
	err = router.Run(config.ApiGatewayHttpHost + config.ApiGatewayHttpPort)
	if err != nil {
		logger.Fatal("Gin router failed to run", zap.Error(err))
		return
	}
}
