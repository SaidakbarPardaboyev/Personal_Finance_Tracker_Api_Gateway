package main

import (
	"api_gateway/api"
	"api_gateway/configs"
	"api_gateway/grpc/client"
	"api_gateway/pkg/logger"
	"api_gateway/storage"
	"context"

	"github.com/casbin/casbin/v2"
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

	storage, err := storage.New(context.Background(), config, &logger)
	if err != nil {
		logger.Fatal("Failed to connect storage ", zap.Error(err))
		return
	}

	casbinEnforcer, err := casbin.NewEnforcer("./configs/auth.conf", "./configs/auth.csv")
	if err != nil {
		logger.Fatal("Failed to create casbin enforcer ", zap.Error(err))
		return
	}

	router := api.NewRouter(&api.Option{
		Log:            logger,
		Services:       services,
		Storage:        storage,
		CasbinEnforcer: casbinEnforcer,
	})

	logger.Info("Gin router is running..")
	err = router.Run(config.ApiGatewayHttpHost + config.ApiGatewayHttpPort)
	if err != nil {
		logger.Fatal("Gin router failed to run", zap.Error(err))
		return
	}
}
