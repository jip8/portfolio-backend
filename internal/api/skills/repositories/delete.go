package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jmoiron/sqlx"
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

func (r *DeleteRepository) Execute(ctx context.Context, ids []int) error {
	query := `DELETE FROM portfolio.skills WHERE id IN (?)`
	executor := r.postgresClient.GetExecutor(ctx)

	query, args, err := sqlx.In(query, ids)
	if err != nil {
		return err
	}

	query = executor.Rebind(query)

	_, err = executor.ExecContext(ctx, query, args...)
	return err
}
