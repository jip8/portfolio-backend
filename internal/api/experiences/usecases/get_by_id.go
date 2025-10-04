package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
)

type GetByIdUC struct {
	config      *entity.Config
	redisClient *redis.Client
}

func NewGetByIdUC(config *entity.Config, redisClient *redis.Client) *GetByIdUC {
	return &GetByIdUC{
		config:      config,
		redisClient: redisClient,
	}
}

func (u *GetByIdUC) Execute(ctx context.Context, id int) (*entity.ExperienceResp, error) {
	

	return nil, nil
}