package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
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

func NewRepository(config *entity.Config, redisClient *redis.Client, postgresClient *services.PostgresClient) articles.Repository {
	return &articlesRepo{
		create:  NewCreateRepository(config, redisClient, postgresClient),
		update:  NewUpdateRepository(config, redisClient, postgresClient),
		delete:  NewDeleteRepository(config, redisClient, postgresClient),
		getById: NewGetByIdRepository(config, redisClient, postgresClient),
		getList: NewGetListRepository(config, redisClient, postgresClient),
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
