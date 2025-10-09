package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/projects"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/links"
)

type CreateUC struct {
	config         *entity.Config
	projectsRepo   projects.Repository
	byId           *GetByIdUC
	postgresClient *services.PostgresClient
	linksUC        links.UseCase
}

func NewCreateUC(config *entity.Config, projectsRepo projects.Repository, byId *GetByIdUC, postgresClient *services.PostgresClient, linksUC links.UseCase) *CreateUC {
	return &CreateUC{
		config:         config,
		projectsRepo:   projectsRepo,
		byId:           byId,
		postgresClient: postgresClient,
		linksUC:        linksUC,
	}
}

func (u *CreateUC) Execute(ctx context.Context, req entity.ProjectFlat) (resp *entity.ProjectResp, err error) {
	err = req.Validate()
	if err != nil {
		return nil, err
	}

	ctx, err = u.postgresClient.StartProcess(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = u.postgresClient.CloseProcess(ctx, err)
	}()

	var createdId *int
	createdId, err = u.projectsRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	err = u.linksUC.Upsert(ctx, req.LinksArray)
	if err != nil {
		return nil, err
	}

	resp, err = u.byId.Execute(ctx, *createdId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
