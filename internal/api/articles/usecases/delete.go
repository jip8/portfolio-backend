package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/articles"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/links"
)

type DeleteUC struct {
	config         *entity.Config
	articlesRepo   articles.Repository
	postgresClient *services.PostgresClient
	linksUC        links.UseCase
}

func NewDeleteUC(config *entity.Config, articlesRepo articles.Repository, postgresClient *services.PostgresClient, linksUC links.UseCase) *DeleteUC {
	return &DeleteUC{
		config:         config,
		articlesRepo:   articlesRepo,
		postgresClient: postgresClient,
		linksUC:        linksUC,
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

	err = u.articlesRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	err = u.linksUC.DeleteAll(ctx, ModuleName, id)
	if err != nil {
		return err
	}

	return nil
}
