package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/projects"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type DeleteUC struct {
	config         *entity.Config
	redisClient    *redis.Client
	projectsRepo   projects.Repository
	postgresClient *services.PostgresClient
}

func NewDeleteUC(config *entity.Config, redisClient *redis.Client, projectsRepo projects.Repository, postgresClient *services.PostgresClient) *DeleteUC {
	return &DeleteUC{
		config:         config,
		redisClient:    redisClient,
		projectsRepo:   projectsRepo,
		postgresClient: postgresClient,
	}
}

func (u *DeleteUC) Execute(ctx context.Context, id int) (err error) {
	ctx, err = u.postgresClient.StartProcess(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = u.postgresClient.CloseProcess(ctx, err)
	}()

	err = u.projectsRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
