package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type DeleteRepository struct {
	config         *entity.Config
	redisClient    *redis.Client
	postgresClient *services.PostgresClient
}

func NewDeleteRepository(config *entity.Config, redisClient *redis.Client, postgresClient *services.PostgresClient) *DeleteRepository {
	return &DeleteRepository{
		config:         config,
		redisClient:    redisClient,
		postgresClient: postgresClient,
	}
}

func (r *DeleteRepository) Execute(ctx context.Context, id int) error {
	query := `DELETE FROM portfolio.projects WHERE id = $1`
	executor := r.postgresClient.GetExecutor(ctx)
	_, err := executor.ExecContext(ctx, query, id)
	return err
}
