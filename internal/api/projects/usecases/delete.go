package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/projects"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/links"
	"github.com/jip/portfolio-backend/internal/api/attachments"
)

type DeleteUC struct {
	config         *entity.Config
	projectsRepo   projects.Repository
	postgresClient *services.PostgresClient
	linksUC        links.UseCase
	attachmentsUC  attachments.UseCase
}

func NewDeleteUC(config *entity.Config, projectsRepo projects.Repository, postgresClient *services.PostgresClient, linksUC links.UseCase, attachmentsUC attachments.UseCase) *DeleteUC {
	return &DeleteUC{
		config:         config,
		projectsRepo:   projectsRepo,
		postgresClient: postgresClient,
		linksUC:        linksUC,
		attachmentsUC:  attachmentsUC,
	}
}

func (u *DeleteUC) Execute(ctx context.Context, id int) (err error) {
	ctx, err = u.postgresClient.StartProcess(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = u.postgresClient.CloseProcess(ctx, err)
	}()

	err = u.projectsRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	err = u.linksUC.DeleteAll(ctx, ModuleName, id)
	if err != nil {
		return err
	}

	err = u.attachmentsUC.DeleteAll(ctx, ModuleName, id)
	if err != nil {
		return err
	}

	return nil
}
