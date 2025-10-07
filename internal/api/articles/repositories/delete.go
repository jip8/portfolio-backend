package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type DeleteRepository struct {
	config         *entity.Config
	postgresClient *services.PostgresClient
}

func NewDeleteRepository(config *entity.Config, postgresClient *services.PostgresClient) *DeleteRepository {
	return &DeleteRepository{
		config:         config,
		postgresClient: postgresClient,
	}
}

func (r *DeleteRepository) Execute(ctx context.Context, id int) error {
	query := `DELETE FROM portfolio.articles WHERE id = $1`
	executor := r.postgresClient.GetExecutor(ctx)
	_, err := executor.ExecContext(ctx, query, id)
	return err
}
