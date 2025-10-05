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

func (r *UpdateRepository) Execute(ctx context.Context, req entity.About) error {
	query := `
		INSERT INTO about_text (id, title, content, updated_at)
        	VALUES (1, :title, :content, NOW())
        ON CONFLICT (id) DO UPDATE SET title = :title, content = EXCLUDED.content, updated_at = NOW();
	`

	executor := r.postgresClient.GetExecutor(ctx)

	namedQuery, args, err := executor.BindNamed(query, req)
	if err != nil {
		return err
	}

	result, err := executor.ExecContext(ctx, namedQuery, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("about not found")
	}

	return nil
}
