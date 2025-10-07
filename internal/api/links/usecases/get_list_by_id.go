package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/links"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetListByIdUC struct {
	config         *entity.Config
	linksRepo      links.Repository
	postgresClient *services.PostgresClient
}

func NewGetListByIdUC(config *entity.Config,linksRepo links.Repository, postgresClient *services.PostgresClient) *GetListByIdUC {
	return &GetListByIdUC{
		config:         config,
		linksRepo:      linksRepo,
		postgresClient: postgresClient,
	}
}

func (u *GetListByIdUC) Execute(ctx context.Context, module string, parent_id int) (entity.LinkRespArray, error) {

	resp, err := u.linksRepo.GetListById(ctx, module, parent_id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
