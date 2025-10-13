package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetListByIdRepository struct {
	config         *entity.Config
	postgresClient *services.PostgresClient
}

func NewGetListByIdRepository(config *entity.Config, postgresClient *services.PostgresClient) *GetListByIdRepository {
	return &GetListByIdRepository{
		config:         config,
		postgresClient: postgresClient,
	}
}

func (r *GetListByIdRepository) Execute(ctx context.Context, module *string, parentId *int) (entity.SkillRespArray, error) {
	var items []entity.SkillResp
	query := `
	SELECT
		t2.id,
		t2.title,
		t2.description
	FROM portfolio.skills_relations t1
	LEFT JOIN portfolio.skills t2 ON t1.skill_id = t2.id
	WHERE module = $1 AND parent_id = $2
	ORDER BY t1.revelance`

	executor := r.postgresClient.GetExecutor(ctx)

	var args []interface{}

	args = append(args, *module, *parentId)

	if err := executor.SelectContext(ctx, &items, query, args...); err != nil {
		return nil, err
	}

	if items == nil {
		items = make([]entity.SkillResp, 0)
	}
	return items, nil
}
