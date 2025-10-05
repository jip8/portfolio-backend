package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetRepository struct {
	config         *entity.Config
	redisClient    *redis.Client
	postgresClient *services.PostgresClient
}

func NewGetRepository(config *entity.Config, redisClient *redis.Client, postgresClient *services.PostgresClient) *GetRepository {
	return &GetRepository{
		config:         config,
		redisClient:    redisClient,
		postgresClient: postgresClient,
	}
}

func (r *GetRepository) Execute(ctx context.Context) (*entity.About, error) {
	query := `
	SELECT
		title,
		content
	FROM portfolio.about WHERE id = 1`

	var about entity.About

	executor := r.postgresClient.GetExecutor(ctx)
	if err := executor.GetContext(ctx, &about, query); err != nil {
		return nil, err
	}

	return &about, nil
}
