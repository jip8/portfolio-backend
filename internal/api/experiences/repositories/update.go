package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"gorm.io/gorm"
)

type UpdateRepository struct {
	config      *entity.Config
	redisClient *redis.Client
	db          *gorm.DB
}

func NewUpdateRepository(config *entity.Config, redisClient *redis.Client, db *gorm.DB) *UpdateRepository {
	return &UpdateRepository{
		config:      config,
		redisClient: redisClient,
		db:          db,
	}
}

func (r *UpdateRepository) Execute(ctx context.Context, req entity.ExperienceFlat) (*int, error) {
	

	return nil, nil
}