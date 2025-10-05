package repositories

import (
	"context"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jmoiron/sqlx"
)

type GetListRepository struct {
	config         *entity.Config
	redisClient    *redis.Client
	postgresClient *services.PostgresClient
}

func NewGetListRepository(config *entity.Config, redisClient *redis.Client, postgresClient *services.PostgresClient) *GetListRepository {
	return &GetListRepository{
		config:         config,
		redisClient:    redisClient,
		postgresClient: postgresClient,
	}
}

func (r *GetListRepository) Execute(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.Contact], error) {
	var items []entity.Contact
	query := `
	SELECT
		id,
		link,
		plataform,
		description,
		revelance
	FROM portfolio.contacts
	WHERE active = TRUE
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
	if err := executor.GetContext(ctx, &total, `SELECT COUNT(*) FROM portfolio.contacts`); err != nil {
		return nil, err
	}

	if len(items) == 0 {
		items = make([]entity.Contact, 0)
	}

	return &entity.List[entity.Contact]{
		Offset: listReq.Offset,
		Limit:  listReq.Limit,
		Total:  total,
		Items:  items,
	}, nil
}
