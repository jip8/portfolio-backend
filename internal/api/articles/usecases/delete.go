package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/articles"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type DeleteUC struct {
	config         *entity.Config
	articlesRepo   articles.Repository
	postgresClient *services.PostgresClient
}

func NewDeleteUC(config *entity.Config, articlesRepo articles.Repository, postgresClient *services.PostgresClient) *DeleteUC {
	return &DeleteUC{
		config:         config,
		articlesRepo:   articlesRepo,
		postgresClient: postgresClient,
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

	return nil
}
