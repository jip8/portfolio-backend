package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/articles"
	"github.com/jip/portfolio-backend/internal/entity"

	"github.com/jip/portfolio-backend"
	"github.com/jip/portfolio-backend/internal/services"
)

type UpdateUC struct {
	config         *entity.Config
	articlesRepo   articles.Repository
	byId           *GetByIdUC
	postgresClient *services.PostgresClient
}

func NewUpdateUC(config *entity.Config, articlesRepo articles.Repository, byId *GetByIdUC, postgresClient *services.PostgresClient) *UpdateUC {
	return &UpdateUC{
		config:         config,
		articlesRepo:   articlesRepo,
		byId:           byId,
		postgresClient: postgresClient,
	}
}

func (u *UpdateUC) Execute(ctx context.Context, req entity.ArticleFlat) (resp *entity.ArticleResp, err error) {
	if req.Id == nil {
		return nil, portfolio.ErrArticleIdIsRequired
	}

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

	var updatedId *int
	updatedId, err = u.articlesRepo.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	resp, err = u.byId.Execute(ctx, *updatedId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
