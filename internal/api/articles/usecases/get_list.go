package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/articles"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetListUC struct {
	config         *entity.Config
	articlesRepo   articles.Repository
	postgresClient *services.PostgresClient
}

func NewGetListUC(config *entity.Config, articlesRepo articles.Repository, postgresClient *services.PostgresClient) *GetListUC {
	return &GetListUC{
		config:         config,
		articlesRepo:   articlesRepo,
		postgresClient: postgresClient,
	}
}

func (u *GetListUC) Execute(ctx context.Context, ListReq entity.ListReq) (*entity.List[entity.ArticleResp], error) {

	resp, err := u.articlesRepo.GetList(ctx, ListReq)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
