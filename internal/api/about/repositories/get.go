package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetRepository struct {
	config         *entity.Config
	postgresClient *services.PostgresClient
}

func NewGetRepository(config *entity.Config, postgresClient *services.PostgresClient) *GetRepository {
	return &GetRepository{
		config:         config,
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
