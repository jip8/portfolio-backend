package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/experiences"
	"github.com/jip/portfolio-backend/internal/entity"
)

type DeleteUC struct {
	config          *entity.Config
	redisClient     *redis.Client
	experiencesRepo experiences.Repository
}

func NewDeleteUC(config *entity.Config, redisClient *redis.Client, experiencesRepo experiences.Repository) *DeleteUC {
	return &DeleteUC{
		config:          config,
		redisClient:     redisClient,
		experiencesRepo: experiencesRepo,
	}
}

func (u *DeleteUC) Execute(ctx context.Context, id int) error {

	err := u.experiencesRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}