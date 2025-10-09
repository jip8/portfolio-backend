package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/articles"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type articlesRepo struct {
	create  *CreateRepository
	update  *UpdateRepository
	delete  *DeleteRepository
	getById *GetByIdRepository
	getList *GetListRepository
}

func NewRepository(config *entity.Config, postgresClient *services.PostgresClient) articles.Repository {
	return &articlesRepo{
		create:  NewCreateRepository(config, postgresClient),
		update:  NewUpdateRepository(config, postgresClient),
		delete:  NewDeleteRepository(config, postgresClient),
		getById: NewGetByIdRepository(config, postgresClient),
		getList: NewGetListRepository(config, postgresClient),
	}
}

func (r *articlesRepo) Create(ctx context.Context, req entity.ArticleFlat) (*int, error) {
	return r.create.Execute(ctx, req)
}

func (r *articlesRepo) Update(ctx context.Context, req entity.ArticleFlat) (*int, error) {
	return r.update.Execute(ctx, req)
}

func (r *articlesRepo) Delete(ctx context.Context, id int) error {
	return r.delete.Execute(ctx, id)
}

func (r *articlesRepo) GetById(ctx context.Context, id int) (*entity.ArticleResp, error) {
	return r.getById.Execute(ctx, id)
}

func (r *articlesRepo) GetList(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.ArticleResp], error) {
	return r.getList.Execute(ctx, listReq)
}
