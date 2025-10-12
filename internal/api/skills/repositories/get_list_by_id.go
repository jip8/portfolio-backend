package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"strings"
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
		id,
		title,
		skill,
		description
	FROM portfolio.skills
	${filter}
	ORDER BY revelance`

	executor := r.postgresClient.GetExecutor(ctx)

	var filter string
	var args []interface{}

	if module != nil && parentId != nil {
		filter = "WHERE module = $1 AND parent_id = $2"
		args = append(args, *module, *parentId)
	} else {
		filter = ""
	}

	query = strings.ReplaceAll(query, "${filter}", filter)

	if err := executor.SelectContext(ctx, &items, query, args...); err != nil {
		return nil, err
	}

	if items == nil {
		items = make([]entity.SkillResp, 0)
	}
	return items, nil
}
