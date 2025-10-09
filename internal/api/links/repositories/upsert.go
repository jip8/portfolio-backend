package repositories

import (
	"context"
	"fmt"
	"strings"

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
	if len(links) == 0 {
		return nil
	}

	valueStrings := make([]string, 0, len(links))
	valueArgs := make([]interface{}, 0, len(links)*6)
	i := 1
	for _, link := range links {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", i, i+1, i+2, i+3, i+4, i+5))
		valueArgs = append(valueArgs, link.ParentId)
		valueArgs = append(valueArgs, link.Module)
		valueArgs = append(valueArgs, link.Title)
		valueArgs = append(valueArgs, link.Link)
		valueArgs = append(valueArgs, link.Revelance)
		valueArgs = append(valueArgs, link.Description)
		i += 6
	}

	stmt := `
        INSERT INTO portfolio.links (parent_id, module, title, link, revelance, description)
        VALUES %s
        ON CONFLICT (parent_id, module, title) DO UPDATE SET
            link = EXCLUDED.link,
            revelance = EXCLUDED.revelance,
            description = EXCLUDED.description,
            updated_at = NOW()
    `
	query := fmt.Sprintf(stmt, strings.Join(valueStrings, ","))

	executor := r.postgresClient.GetExecutor(ctx)
	_, err := executor.ExecContext(ctx, query, valueArgs...)
	return err
}
