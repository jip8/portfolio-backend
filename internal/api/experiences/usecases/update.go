package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
)

type UpdateUC struct {
	config      *entity.Config
	redisClient *redis.Client
}

func NewUpdateUC(config *entity.Config, redisClient *redis.Client) *UpdateUC {
	return &UpdateUC{
		config:      config,
		redisClient: redisClient,
	}
}

func (u *UpdateUC) Execute(ctx context.Context, req entity.ExperienceFlat) (*entity.ExperienceResp, error) {
	

	return nil, nil
}