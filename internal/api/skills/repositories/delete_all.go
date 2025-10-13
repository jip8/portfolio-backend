package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type DeleteAllRepository struct {
	config         *entity.Config
	postgresClient *services.PostgresClient
}

func NewDeleteAllRepository(config *entity.Config, postgresClient *services.PostgresClient) *DeleteAllRepository {
	return &DeleteAllRepository{
		config:         config,
		postgresClient: postgresClient,
	}
}

func (r *DeleteAllRepository) Execute(ctx context.Context, module string, parent_id int) error {
	query := `DELETE FROM portfolio.skills_relations WHERE module = $1 AND parent_id = $2`
	executor := r.postgresClient.GetExecutor(ctx)

	_, err := executor.ExecContext(ctx, query, module, parent_id)
	return err
}
