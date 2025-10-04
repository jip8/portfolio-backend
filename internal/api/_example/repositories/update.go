package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UpdateRepository struct {
	config      *entity.Config
	redisClient *redis.Client
	db          *sqlx.DB
}

func NewUpdateRepository(config *entity.Config, redisClient *redis.Client, db *sqlx.DB) *UpdateRepository {
	return &UpdateRepository{
		config:      config,
		redisClient: redisClient,
		db:          db,
	}
}

func (r *UpdateRepository) Execute(ctx context.Context, req entity.ExperienceFlat) (*entity.ExperienceResp, error) {
	

	return nil, nil
}