package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetByIdRepository struct {
	config         *entity.Config
	postgresClient *services.PostgresClient
}

func NewGetByIdRepository(config *entity.Config, postgresClient *services.PostgresClient) *GetByIdRepository {
	return &GetByIdRepository{
		config:         config,
		postgresClient: postgresClient,
	}
}

func (r *GetByIdRepository) Execute(ctx context.Context, id int) (*entity.CourseResp, error) {
	query := `
	SELECT
		id,
		title,
		description,
		concluded_at AS concluded_at_time,
		revelance
	FROM portfolio.courses WHERE id = $1`

	var course entity.CourseResp

	executor := r.postgresClient.GetExecutor(ctx)
	if err := executor.GetContext(ctx, &course, query, id); err != nil {
		return nil, err
	}

	return &course, nil
}
