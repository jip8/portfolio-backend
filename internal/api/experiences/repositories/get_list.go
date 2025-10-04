package repositories

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetListRepository struct {
	config      *entity.Config
	redisClient *redis.Client
	db          *sqlx.DB
}

func NewGetListRepository(config *entity.Config, redisClient *redis.Client, db *sqlx.DB) *GetListRepository {
	return &GetListRepository{
		config:      config,
		redisClient: redisClient,
		db:          db,
	}
}

func (r *GetListRepository) Execute(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.ExperienceResp], error) {
	var items []entity.Experience
	query := `SELECT * FROM experiences LIMIT $1 OFFSET $2`

	if err := r.db.SelectContext(ctx, &items, query, *listReq.Limit, *listReq.Offset); err != nil {
		return nil, err
	}

	var total int
	if err := r.db.GetContext(ctx, &total, `SELECT COUNT(*) FROM experiences`); err != nil {
		return nil, err
	}

	respItems := make([]entity.ExperienceResp, len(items))
	for i, item := range items {
		initialDate := item.InitialDate.Format("2006-01-02")
		endDate := item.EndDate.Format("2006-01-02")
		respItems[i] = entity.ExperienceResp{
			Id:          strconv.Itoa(item.ID),
			Title:       item.Title,
			Function:    &item.Function,
			Description: &item.Description,
			InitialDate: &initialDate,
			EndDate:     &endDate,
		}
	}

	return &entity.List[entity.ExperienceResp]{
		Offset: *listReq.Offset,
		Limit:  *listReq.Limit,
		Total:  total,
		Items:  respItems,
	}, nil
}