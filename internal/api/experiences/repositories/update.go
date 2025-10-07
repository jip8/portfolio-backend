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

func (r *UpdateRepository) Execute(ctx context.Context, req entity.ExperienceFlat) (*int, error) {
	query := `
		UPDATE portfolio.experiences
		SET
			title = :title,
			"function" = :function,
			description = :description,
			initial_date = :initial_date_time,
			end_date = :end_date_time,
			actual = :actual,
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
		return nil, errors.New("experience not found")
	}

	return req.Id, nil
}