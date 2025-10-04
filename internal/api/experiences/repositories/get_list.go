package repositories

import (
	"context"

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

func (r *GetListRepository) Execute(ctx context.Context, listReq entity.ListReq) (*entity.List, error) {
	

	return nil, nil
}