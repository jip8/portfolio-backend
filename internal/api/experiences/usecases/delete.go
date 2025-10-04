package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
)

type DeleteUC struct {
	config      *entity.Config
	redisClient *redis.Client
}

func NewDeleteUC(config *entity.Config, redisClient *redis.Client) *DeleteUC {
	return &DeleteUC{
		config:      config,
		redisClient: redisClient,
	}
}

func (u *DeleteUC) Execute(ctx context.Context, id int) error {
	

	return nil
}