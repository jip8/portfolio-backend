package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/about"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type aboutUC struct {
	update  *UpdateUC
	get 	*GetUC
}

func NewUseCase(config *entity.Config, redisClient *redis.Client, aboutRepo about.Repository, postgresClient *services.PostgresClient) about.UseCase {
	getUC := NewGetUC(config, redisClient, aboutRepo, postgresClient)

	return &aboutUC{
		update:  	NewUpdateUC(config, redisClient, aboutRepo, getUC, postgresClient),
		get: 		getUC,
	}
}

func (u *aboutUC) Update(ctx context.Context, req entity.About) (*entity.About, error) {
	return u.update.Execute(ctx, req)
}

func (u *aboutUC) Get(ctx context.Context) (*entity.About, error) {
	return u.get.Execute(ctx)
}
