package storage

import (
	"api_gateway/configs"
	"api_gateway/pkg/logger"
	rds "api_gateway/storage/redis"
	"context"

	"github.com/redis/go-redis/v9"
)

type Storage struct {
	redisClient *redis.Client
	log         logger.ILogger
}

type IStorage interface {
	Close()
	RedisClient() IUserRedisStorage
}

type IUserRedisStorage interface {
	SaveCodeWithEmail(context.Context, string, string) error
	GetCodeWithEmail(context.Context, string) (string, error)
}

func New(ctx context.Context, cfg *configs.Config, log *logger.ILogger) (IStorage, error) {
	redisClient, err := rds.NewRedisClient()
	if err != nil {
		return nil, err
	}

	return &Storage{
		redisClient: redisClient,
		log:         *log,
	}, nil
}

func (s *Storage) Close() {
	s.redisClient.Close()
}

func (s *Storage) RedisClient() IUserRedisStorage {
	return rds.NewUsersRedisRepo(s.redisClient, s.log)
}
