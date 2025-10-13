package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type InsertRepository struct {
	config         *entity.Config
	postgresClient *services.PostgresClient
}

func NewInsertRepository(config *entity.Config, postgresClient *services.PostgresClient) *InsertRepository {
	return &InsertRepository{
		config:         config,
		postgresClient: postgresClient,
	}
}

func (r *InsertRepository) Execute(ctx context.Context, attachments []entity.AttachmentFlat) error {
	if len(attachments) == 0 {
		return nil
	}

	valueStrings := make([]string, 0, len(attachments))
	valueArgs := make([]interface{}, 0, len(attachments)*7)
	i := 1
	for _, attachment := range attachments {
		if attachment.Id == nil || *attachment.Id == 0 {
			valueStrings = append(valueStrings,
				fmt.Sprintf("(DEFAULT, $%d, $%d, $%d, $%d, $%d, $%d)",
					i, i+1, i+2, i+3, i+4, i+5))
			valueArgs = append(valueArgs,
				attachment.ParentId,
				attachment.Module,
				attachment.Title,
				attachment.Link,
				attachment.ContentType,
				attachment.Description,
			)
			i += 6
		} else {
			valueStrings = append(valueStrings,
				fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)",
					i, i+1, i+2, i+3, i+4, i+5, i+6))
			valueArgs = append(valueArgs,
				*attachment.Id,
				attachment.ParentId,
				attachment.Module,
				attachment.Title,
				attachment.Link,
				attachment.ContentType,
				attachment.Description,
			)
			i += 7
		}
	}

	stmt := `
        INSERT INTO portfolio.attachments (id, parent_id, module, title, link, content_type, description)
        VALUES %s
        ON CONFLICT (id) DO UPDATE SET
            parent_id = EXCLUDED.parent_id,
            module = EXCLUDED.module,
            title = EXCLUDED.title,
            link = EXCLUDED.link,
            content_type = EXCLUDED.content_type,
            description = EXCLUDED.description,
            updated_at = NOW()
    `
	query := fmt.Sprintf(stmt, strings.Join(valueStrings, ","))

	executor := r.postgresClient.GetExecutor(ctx)
	_, err := executor.ExecContext(ctx, query, valueArgs...)
	return err
}