package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GetByIdRepository struct {
	config      *entity.Config
	redisClient *redis.Client
	db          *sqlx.DB
}

func NewGetByIdRepository(config *entity.Config, redisClient *redis.Client, db *sqlx.DB) *GetByIdRepository {
	return &GetByIdRepository{
		config:      config,
		redisClient: redisClient,
		db:          db,
	}
}

func (r *GetByIdRepository) Execute(ctx context.Context, id int) (*entity.ExperienceResp, error) {
	

	return nil, nil
}