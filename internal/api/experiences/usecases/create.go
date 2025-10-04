package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
)

type CreateUC struct {
	config      *entity.Config
	redisClient *redis.Client
}

func NewCreateUC(config *entity.Config, redisClient *redis.Client) *CreateUC {
	return &CreateUC{
		config:      config,
		redisClient: redisClient,
	}
}

func (u *CreateUC) Execute(ctx context.Context, req entity.ExperienceFlat) (*entity.ExperienceResp, error) {
	

	return nil, nil
}