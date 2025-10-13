package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetListRepository struct {
	config         *entity.Config
	postgresClient *services.PostgresClient
}

func NewGetListRepository(config *entity.Config, postgresClient *services.PostgresClient) *GetListRepository {
	return &GetListRepository{
		config:         config,
		postgresClient: postgresClient,
	}
}

func (r *GetListRepository) Execute(ctx context.Context) (entity.SkillRespArray, error) {
	var items []entity.SkillResp
	query := `
	SELECT
		id,
		title,
		description
	FROM portfolio.skills
	ORDER BY id`

	executor := r.postgresClient.GetExecutor(ctx)

	var args []interface{}

	if err := executor.SelectContext(ctx, &items, query, args...); err != nil {
		return nil, err
	}

	if items == nil {
		items = make([]entity.SkillResp, 0)
	}
	return items, nil
}
