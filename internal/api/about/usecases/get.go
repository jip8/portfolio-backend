package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/about"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetUC struct {
	config         *entity.Config
	aboutRepo      about.Repository
	postgresClient *services.PostgresClient
}

func NewGetUC(config *entity.Config, aboutRepo about.Repository, postgresClient *services.PostgresClient) *GetUC {
	return &GetUC{
		config:         config,
		aboutRepo:      aboutRepo,
		postgresClient: postgresClient,
	}
}

func (u *GetUC) Execute(ctx context.Context) (*entity.About, error) {

	resp, err := u.aboutRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
