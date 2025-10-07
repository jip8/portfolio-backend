package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type UpsertRepository struct {
	config         *entity.Config
	postgresClient *services.PostgresClient
}

func NewUpsertRepository(config *entity.Config, postgresClient *services.PostgresClient) *UpsertRepository {
	return &UpsertRepository{
		config:         config,
		postgresClient: postgresClient,
	}
}

func (r *UpsertRepository) Execute(ctx context.Context, links []entity.LinkFlat) error {
	query := `
        INSERT INTO portfolio.links (parent_id, module, title, link, revelance, description)
        VALUES (:parent_id, :module, :title, :link, :revelance, :description)
        ON CONFLICT (parent_id, module, title) DO UPDATE SET
            link = EXCLUDED.link,
            revelance = EXCLUDED.revelance,
            description = EXCLUDED.description,
            updated_at = NOW()
    `

	executor := r.postgresClient.GetExecutor(ctx)

	_, err := executor.NamedExecContext(ctx, query, links)
	return err
}