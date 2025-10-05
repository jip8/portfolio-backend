package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type CreateRepository struct {
	config         *entity.Config
	redisClient    *redis.Client
	postgresClient *services.PostgresClient
}

func NewCreateRepository(config *entity.Config, redisClient *redis.Client, postgresClient *services.PostgresClient) *CreateRepository {
	return &CreateRepository{
		config:         config,
		redisClient:    redisClient,
		postgresClient: postgresClient,
	}
}

func (r *CreateRepository) Execute(ctx context.Context, req entity.ArticleFlat) (*int, error) {
	query := `
		INSERT INTO portfolio.articles 
			(type, title, description, local, published_at, revelance)
		VALUES 
			(:type, :title, :description, :local, :published_at_time, :revelance)
		RETURNING id
	`

	executor := r.postgresClient.GetExecutor(ctx)

	namedQuery, args, err := executor.BindNamed(query, req)
	if err != nil {
		return nil, err
	}

	var newID int
	err = executor.QueryRowContext(ctx, namedQuery, args...).Scan(&newID)
	if err != nil {
		return nil, err
	}

	return &newID, nil
}