package repositories

import (
	"context"
	"errors"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type UpdateRepository struct {
	config         *entity.Config
	postgresClient *services.PostgresClient
}

func NewUpdateRepository(config *entity.Config, postgresClient *services.PostgresClient) *UpdateRepository {
	return &UpdateRepository{
		config:         config,
		postgresClient: postgresClient,
	}
}

func (r *UpdateRepository) Execute(ctx context.Context, req entity.ArticleFlat) (*int, error) {
	query := `
		UPDATE portfolio.articles
		SET
			type = :type,
			title = :title,
			description = :description,
			local = :local,
			published_at = :published_at_time,
			revelance = :revelance,
			thumbnail_id = :thumbnail_id,
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
		return nil, errors.New("article not found")
	}

	return req.Id, nil
}
