package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/projects"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetByIdUC struct {
	config         *entity.Config
	redisClient    *redis.Client
	projectsRepo   projects.Repository
	postgresClient *services.PostgresClient
}

func NewGetByIdUC(config *entity.Config, redisClient *redis.Client, projectsRepo projects.Repository, postgresClient *services.PostgresClient) *GetByIdUC {
	return &GetByIdUC{
		config:         config,
		redisClient:    redisClient,
		projectsRepo:   projectsRepo,
		postgresClient: postgresClient,
	}
}

func (u *GetByIdUC) Execute(ctx context.Context, id int) (*entity.ProjectResp, error) {

	resp, err := u.projectsRepo.GetById(ctx, id)
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
