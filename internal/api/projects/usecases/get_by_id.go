package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/projects"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/links"
)

type GetByIdUC struct {
	config         *entity.Config
	projectsRepo   projects.Repository
	postgresClient *services.PostgresClient
	linksUC        links.UseCase
}

func NewGetByIdUC(config *entity.Config, projectsRepo projects.Repository, postgresClient *services.PostgresClient, linksUC links.UseCase) *GetByIdUC {
	return &GetByIdUC{
		config:         config,
		projectsRepo:   projectsRepo,
		postgresClient: postgresClient,
		linksUC:        linksUC,
	}
}

func (u *GetByIdUC) Execute(ctx context.Context, id int) (*entity.ProjectResp, error) {

	resp, err := u.projectsRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		err = resp.Format()
		if err != nil {
			return nil, err
		}
	}

	links, err := u.linksUC.GetListById(ctx, ModuleName, id)
	if err != nil {
		return nil, err
	}

	resp.LinksRespArray = links

	return resp, nil
}
