package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/api/experiences"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetByIdUC struct {
	config      	*entity.Config
	redisClient 	*redis.Client
	experiencesRepo experiences.Repository
	postgresClient 	*services.PostgresClient
}

func NewGetByIdUC(config *entity.Config, redisClient *redis.Client, experiencesRepo experiences.Repository, postgresClient *services.PostgresClient) *GetByIdUC {
	return &GetByIdUC{
		config:      		config,
		redisClient: 		redisClient,
		experiencesRepo: 	experiencesRepo,
		postgresClient: 	postgresClient,
	}
}

func (u *GetByIdUC) Execute(ctx context.Context, id int) (*entity.ExperienceResp, error) {
	
	resp, err := u.experiencesRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		err = resp.Format()
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}