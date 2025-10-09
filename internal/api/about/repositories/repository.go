package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/about"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type aboutRepo struct {
	update *UpdateRepository
	get    *GetRepository
}

func NewRepository(config *entity.Config, postgresClient *services.PostgresClient) about.Repository {
	return &aboutRepo{
		update: NewUpdateRepository(config, postgresClient),
		get:    NewGetRepository(config, postgresClient),
	}
}

func (r *aboutRepo) Update(ctx context.Context, req entity.About) error {
	return r.update.Execute(ctx, req)
}

func (r *aboutRepo) Get(ctx context.Context) (*entity.About, error) {
	return r.get.Execute(ctx)
}
