package repositories

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type UpdateRepository struct {
	config         *entity.Config
	redisClient    *redis.Client
	postgresClient *services.PostgresClient
}

func NewUpdateRepository(config *entity.Config, redisClient *redis.Client, postgresClient *services.PostgresClient) *UpdateRepository {
	return &UpdateRepository{
		config:         config,
		redisClient:    redisClient,
		postgresClient: postgresClient,
	}
}

func (r *UpdateRepository) Execute(ctx context.Context, req entity.Contact) (*int, error) {
	query := `
		UPDATE portfolio.contacts
		SET
			link = :link,
			plataform = :plataform,
			description = :description,
			active = :active,
			revelance = :revelance,
			updated_at = NOW()
		WHERE
			id = :id
	`

	executor := r.postgresClient.GetExecutor(ctx)

	namedQuery, args, err := executor.BindNamed(query, req)
	if err != nil {
		return nil, err
	}

	result, err := executor.ExecContext(ctx, namedQuery, args...)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("contact not found")
	}

	return req.Id, nil
}

