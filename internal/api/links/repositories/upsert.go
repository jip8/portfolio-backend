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
	valueArgs := make([]interface{}, 0, len(links)*7)
	i := 1

	for _, link := range links {
		if link.Id == nil || *link.Id == 0 {
			valueStrings = append(valueStrings,
				fmt.Sprintf("(DEFAULT, $%d, $%d, $%d, $%d, $%d, $%d)", i, i+1, i+2, i+3, i+4, i+5))

			valueArgs = append(valueArgs,
				link.ParentId,
				link.Module,
				link.Title,
				link.Link,
				link.Revelance,
				link.Description,
			)
			i += 6
		} else {
			valueStrings = append(valueStrings,
				fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)", i, i+1, i+2, i+3, i+4, i+5, i+6))

			valueArgs = append(valueArgs,
				*link.Id,
				link.ParentId,
				link.Module,
				link.Title,
				link.Link,
				link.Revelance,
				link.Description,
			)
			i += 7
		}
	}

	stmt := `
        INSERT INTO portfolio.links (id, parent_id, module, title, link, revelance, description)
        VALUES %s
        ON CONFLICT (id) DO UPDATE SET
            parent_id = EXCLUDED.parent_id,
            module = EXCLUDED.module,
            title = EXCLUDED.title,
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
