package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/courses"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type DeleteUC struct {
	config         *entity.Config
	redisClient    *redis.Client
	coursesRepo    courses.Repository
	postgresClient *services.PostgresClient
}

func NewDeleteUC(config *entity.Config, redisClient *redis.Client, coursesRepo courses.Repository, postgresClient *services.PostgresClient) *DeleteUC {
	return &DeleteUC{
		config:         config,
		redisClient:    redisClient,
		coursesRepo:    coursesRepo,
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

	err = u.coursesRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
