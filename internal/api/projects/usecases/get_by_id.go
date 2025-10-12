package usecases

import (
	"fmt"
	"context"

	"github.com/jip/portfolio-backend/internal/api/projects"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/links"
	"github.com/jip/portfolio-backend/internal/api/attachments"
	"github.com/jip/portfolio-backend/internal/api/skills"
)

type GetByIdUC struct {
	config         *entity.Config
	projectsRepo   projects.Repository
	postgresClient *services.PostgresClient
	linksUC        links.UseCase
	attachmentsUC  attachments.UseCase
	skillsUC       skills.UseCase
}

func NewGetByIdUC(config *entity.Config, projectsRepo projects.Repository, postgresClient *services.PostgresClient, linksUC links.UseCase, attachmentUC attachments.UseCase, skillsUC skills.UseCase) *GetByIdUC {
	return &GetByIdUC{
		config:         config,
		projectsRepo:   projectsRepo,
		postgresClient: postgresClient,
		linksUC:        linksUC,
		attachmentsUC:  attachmentUC,
		skillsUC:       skillsUC,
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

	attachments, err := u.attachmentsUC.GetListById(ctx, ModuleName, id)
	if err != nil {
		return nil, err
	}
	
	resp.Attachments = &attachments

	module := fmt.Sprintf("%s", ModuleName)
	resp.Skills, err = u.skillsUC.GetList(ctx, &module, &id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
