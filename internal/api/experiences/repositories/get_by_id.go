package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
	"gorm.io/gorm"
)

type GetByIdRepository struct {
	config      *entity.Config
	redisClient *redis.Client
	db          *gorm.DB
}

func NewGetByIdRepository(config *entity.Config, redisClient *redis.Client, db *gorm.DB) *GetByIdRepository {
	return &GetByIdRepository{
		config:      config,
		redisClient: redisClient,
		db:          db,
	}
}

func (r *GetByIdRepository) Execute(ctx context.Context, id int) (*entity.ExperienceResp, error) {
	

	return nil, nil
}