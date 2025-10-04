package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"gorm.io/gorm"
)

type CreateRepository struct {
	config      *entity.Config
	redisClient *redis.Client
	db          *gorm.DB
}

func NewCreateRepository(config *entity.Config, redisClient *redis.Client, db *gorm.DB) *CreateRepository {
	return &CreateRepository{
		config:      config,
		redisClient: redisClient,
		db:          db,
	}
}

func (r *CreateRepository) Execute(ctx context.Context, req entity.ExperienceFlat) (*int, error) {
	

	return nil, nil
}