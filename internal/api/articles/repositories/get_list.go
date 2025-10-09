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

func NewGetListRepository(config *entity.Config,  postgresClient *services.PostgresClient) *GetListRepository {
	return &GetListRepository{
		config:         config,
		postgresClient: postgresClient,
	}
}

func (r *GetListRepository) Execute(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.ArticleResp], error) {
	var items []entity.ArticleResp
	query := `
	SELECT
		id,
		type,
		title,
		description,
		local,
		published_at AS published_at_time,
		revelance
	FROM portfolio.articles
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
	if err := executor.GetContext(ctx, &total, `SELECT COUNT(*) FROM portfolio.articles`); err != nil {
		return nil, err
	}

	if len(items) > 0 {
		for i := range items {
			items[i].Format()
		}
	} else {
		items = make([]entity.ArticleResp, 0)
	}

	return &entity.List[entity.ArticleResp]{
		Offset: listReq.Offset,
		Limit:  listReq.Limit,
		Total:  total,
		Items:  items,
	}, nil
}