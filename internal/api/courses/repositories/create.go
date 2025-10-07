package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type CreateRepository struct {
	config         *entity.Config
	postgresClient *services.PostgresClient
}

func NewCreateRepository(config *entity.Config, postgresClient *services.PostgresClient) *CreateRepository {
	return &CreateRepository{
		config:         config,
		postgresClient: postgresClient,
	}
}

func (r *CreateRepository) Execute(ctx context.Context, req entity.CourseFlat) (*int, error) {
	query := `
		INSERT INTO portfolio.courses 
			(title, description, concluded_at, revelance)
		VALUES 
			(:title, :description, :concluded_at_time, :revelance)
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

