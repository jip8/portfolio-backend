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

func (r *GetByIdRepository) Execute(ctx context.Context, id int) (*entity.Contact, error) {
	query := `
	SELECT
		id,
		link,
		plataform,
		description,
		active,
		revelance
	FROM portfolio.contacts WHERE id = $1`

	var contact entity.Contact

	executor := r.postgresClient.GetExecutor(ctx)
	if err := executor.GetContext(ctx, &contact, query, id); err != nil {
		return nil, err
	}

	return &contact, nil
}
