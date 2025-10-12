package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/articles"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/links"
	"github.com/jip/portfolio-backend/internal/api/attachments"
)

var (
	ModuleName = "articles"
)

type articlesUC struct {
	create  *CreateUC
	update  *UpdateUC
	delete  *DeleteUC
	getById *GetByIdUC
	getList *GetListUC
}

func NewUseCase(config *entity.Config,  articlesRepo articles.Repository, postgresClient *services.PostgresClient, linksUC links.UseCase, attachmentsUC attachments.UseCase) articles.UseCase {
	byId := NewGetByIdUC(config, articlesRepo, postgresClient, linksUC, attachmentsUC)

	return &articlesUC{
		create:  NewCreateUC(config, articlesRepo, byId, postgresClient, linksUC),
		update:  NewUpdateUC(config, articlesRepo, byId, postgresClient, linksUC),
		delete:  NewDeleteUC(config, articlesRepo, postgresClient, linksUC, attachmentsUC),
		getById: byId,
		getList: NewGetListUC(config, articlesRepo, postgresClient),
	}
}

func (u *articlesUC) Create(ctx context.Context, req entity.ArticleFlat) (*entity.ArticleResp, error) {
	return u.create.Execute(ctx, req)
}

func (u *articlesUC) Update(ctx context.Context, req entity.ArticleFlat) (*entity.ArticleResp, error) {
	return u.update.Execute(ctx, req)
}

func (u *articlesUC) Delete(ctx context.Context, id int) error {
	return u.delete.Execute(ctx, id)
}

func (u *articlesUC) GetById(ctx context.Context, id int) (*entity.ArticleResp, error) {
	return u.getById.Execute(ctx, id)
}

func (u *articlesUC) GetList(ctx context.Context, ListReq entity.ListReq) (*entity.List[entity.ArticleResp], error) {
	return u.getList.Execute(ctx, ListReq)
}
