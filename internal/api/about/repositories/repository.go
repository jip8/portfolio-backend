package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/about"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type aboutRepo struct {
	update *UpdateRepository
	get    *GetRepository
}

func NewRepository(config *entity.Config, redisClient *redis.Client, postgresClient *services.PostgresClient) about.Repository {
	return &aboutRepo{
		update: NewUpdateRepository(config, redisClient, postgresClient),
		get:    NewGetRepository(config, redisClient, postgresClient),
	}
}

func (r *aboutRepo) Update(ctx context.Context, req entity.About) error {
	return r.update.Execute(ctx, req)
}

func (r *aboutRepo) Get(ctx context.Context) (*entity.About, error) {
	return r.get.Execute(ctx)
}
