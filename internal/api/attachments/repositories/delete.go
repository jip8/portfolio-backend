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

func (r *DeleteRepository) Execute(ctx context.Context, module string, parentId int, ids []int) error {
	if len(ids) == 0 {
		return nil
	}
	query := `DELETE FROM portfolio.attachments WHERE module = ? AND parent_id = ? AND id IN (?)`

	query, args, err := sqlx.In(query, module, parentId, ids)
	if err != nil {
		return err
	}

	executor := r.postgresClient.GetExecutor(ctx)
	query = executor.Rebind(query)

	_, err = executor.ExecContext(ctx, query, args...)
	return err
}