package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/articles"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/links"
)

type GetByIdUC struct {
	config         *entity.Config
	articlesRepo   articles.Repository
	postgresClient *services.PostgresClient
	linksUC        links.UseCase

}

func NewGetByIdUC(config *entity.Config, articlesRepo articles.Repository, postgresClient *services.PostgresClient, linksUC links.UseCase) *GetByIdUC {
	return &GetByIdUC{
		config:         config,
		articlesRepo:   articlesRepo,
		postgresClient: postgresClient,
		linksUC:        linksUC,
	}
}

func (u *GetByIdUC) Execute(ctx context.Context, id int) (*entity.ArticleResp, error) {

	resp, err := u.articlesRepo.GetById(ctx, id)
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
