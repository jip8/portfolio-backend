package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CreateRepository struct {
	config      *entity.Config
	redisClient *redis.Client
	db          *sqlx.DB
}

func NewCreateRepository(config *entity.Config, redisClient *redis.Client, db *sqlx.DB) *CreateRepository {
	return &CreateRepository{
		config:      config,
		redisClient: redisClient,
		db:          db,
	}
}

func (r *CreateRepository) Execute(ctx context.Context, req entity.ExperienceFlat) (*entity.ExperienceResp, error) {
	

	return nil, nil
}