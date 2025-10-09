package repositories

import (
	"context"
	"strings"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jmoiron/sqlx"
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

func (r *GetListRepository) Execute(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.ExperienceResp], error) {
	var items []entity.ExperienceResp
	query := `
	SELECT
		id,
		title,
		"function",
		description,
		initial_date AS initial_date_time,
		end_date AS end_date_time,
		actual
	FROM portfolio.experiences
	ORDER BY ${order}
	LIMIT :limit
	OFFSET :offset`

	query = strings.ReplaceAll(query, "${order}", listReq.Order)

	namedQuery, args, err := sqlx.Named(query, listReq)
	if err != nil {
		return nil, err
	}

	executor := r.postgresClient.GetExecutor(ctx)

	namedQuery = executor.Rebind(namedQuery)

	if err := executor.SelectContext(ctx, &items, namedQuery, args...); err != nil {
		return nil, err
	}

	var total int
	if err := executor.GetContext(ctx, &total, `SELECT COUNT(*) FROM portfolio.experiences`); err != nil {
		return nil, err
	}

	if len(items) > 0 {
		for i := range items {
			items[i].Format()
		}
	} else {
		items = make([]entity.ExperienceResp, 0)
	}

	return &entity.List[entity.ExperienceResp]{
		Offset: listReq.Offset,
		Limit:  listReq.Limit,
		Total:  total,
		Items:  items,
	}, nil
}