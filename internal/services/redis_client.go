package services

import (
	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
)

func NewRedisClient(cfg *entity.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	return rdb, nil
}
