package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type AddExclusiveRepository struct {
	config         *entity.Config
	postgresClient *services.PostgresClient
}

func NewAddExclusiveRepository(config *entity.Config, postgresClient *services.PostgresClient) *AddExclusiveRepository {
	return &AddExclusiveRepository{
		config:         config,
		postgresClient: postgresClient,
	}
}

func (r *AddExclusiveRepository) Execute(ctx context.Context, parent_id *int, module *string, ids []int) error {
	var err error
	
	if parent_id == nil || module == nil {
		return nil
	}

	err = r.DeleteAll(ctx, *module, *parent_id)
	if err != nil {
		return err
	}

	if len(ids) > 0 {
		query := `
		INSERT INTO portfolio.skills_relations (parent_id, module, skill_id)
		VALUES %s
		`
		valueStrings := make([]string, 0, len(ids))
		valueArgs := make([]interface{}, 0, len(ids))
		i := 1
		for _, id := range ids {
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i, i+1, i+2))
			valueArgs = append(valueArgs, parent_id)
			valueArgs = append(valueArgs, module)
			valueArgs = append(valueArgs, id)
			i += 3
		}
		
		query = fmt.Sprintf(query, strings.Join(valueStrings, ","))
		
		executor := r.postgresClient.GetExecutor(ctx)
		_, err = executor.ExecContext(ctx, query, valueArgs...)
	}

	return err
}

func (r *AddExclusiveRepository) DeleteAll(ctx context.Context, module string, parent_id int) error {
	query := `DELETE FROM portfolio.skills_relations WHERE module = $1 AND parent_id = $2`
	executor := r.postgresClient.GetExecutor(ctx)

	_, err := executor.ExecContext(ctx, query, module, parent_id)
	return err
}