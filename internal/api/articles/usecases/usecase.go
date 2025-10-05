package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/articles"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type articlesUC struct {
	create  *CreateUC
	update  *UpdateUC
	delete  *DeleteUC
	getById *GetByIdUC
	getList *GetListUC
}

func NewUseCase(config *entity.Config, redisClient *redis.Client, articlesRepo articles.Repository, postgresClient *services.PostgresClient) articles.UseCase {
	byId := NewGetByIdUC(config, redisClient, articlesRepo, postgresClient)

	return &articlesUC{
		create:  NewCreateUC(config, redisClient, articlesRepo, byId, postgresClient),
		update:  NewUpdateUC(config, redisClient, articlesRepo, byId, postgresClient),
		delete:  NewDeleteUC(config, redisClient, articlesRepo, postgresClient),
		getById: byId,
		getList: NewGetListUC(config, redisClient, articlesRepo, postgresClient),
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
