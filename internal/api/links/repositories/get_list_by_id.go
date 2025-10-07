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

func (r *GetListByIdRepository) Execute(ctx context.Context, module string, parentId int) (entity.LinkRespArray, error) {
	var items []entity.LinkResp
	query := `
	SELECT
		id,
		title,
		link,
		description
	FROM portfolio.links
	WHERE module = $1 AND parent_id = $2
	ORDER BY revelance`

	executor := r.postgresClient.GetExecutor(ctx)

	if err := executor.SelectContext(ctx, &items, query, module, parentId); err != nil {
		return nil, err
	}

	if items == nil {
		items = make([]entity.LinkResp, 0)
	}

	return items, nil
}