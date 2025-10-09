package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/about"
	"github.com/jip/portfolio-backend/internal/entity"

	"github.com/jip/portfolio-backend/internal/services"
)

type UpdateUC struct {
	config         *entity.Config
	aboutRepo      about.Repository
	get           *GetUC
	postgresClient *services.PostgresClient
}

func NewUpdateUC(config *entity.Config, aboutRepo about.Repository, get *GetUC, postgresClient *services.PostgresClient) *UpdateUC {
	return &UpdateUC{
		config:         config,
		aboutRepo:      aboutRepo,
		get:           get,
		postgresClient: postgresClient,
	}
}

func (u *UpdateUC) Execute(ctx context.Context, req entity.About) (resp *entity.About, err error) {

	ctx, err = u.postgresClient.StartProcess(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = u.postgresClient.CloseProcess(ctx, err)
	}()

	err = u.aboutRepo.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	resp, err = u.get.Execute(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
