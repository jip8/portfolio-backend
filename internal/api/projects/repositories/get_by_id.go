package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetByIdRepository struct {
	config         *entity.Config
	redisClient    *redis.Client
	postgresClient *services.PostgresClient
}

func NewGetByIdRepository(config *entity.Config, redisClient *redis.Client, postgresClient *services.PostgresClient) *GetByIdRepository {
	return &GetByIdRepository{
		config:         config,
		redisClient:    redisClient,
		postgresClient: postgresClient,
	}
}

func (r *GetByIdRepository) Execute(ctx context.Context, id int) (*entity.ProjectResp, error) {
	query := `
	SELECT
		id,
		title,
		description,
		published_at AS published_at_time,
		revelance
	FROM portfolio.projects WHERE id = $1`

	var project entity.ProjectResp

	executor := r.postgresClient.GetExecutor(ctx)
	if err := executor.GetContext(ctx, &project, query, id); err != nil {
		return nil, err
	}

	return &project, nil
}
