package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"gorm.io/gorm"
)

type DeleteRepository struct {
	config      *entity.Config
	redisClient *redis.Client
	db          *gorm.DB
}

func NewDeleteRepository(config *entity.Config, redisClient *redis.Client, db *gorm.DB) *DeleteRepository {
	return &DeleteRepository{
		config:      config,
		redisClient: redisClient,
		db:          db,
	}
}

func (r *DeleteRepository) Execute(ctx context.Context, id int) error {
	
	

	return nil
}