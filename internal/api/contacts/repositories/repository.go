package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/contacts"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type contactsRepo struct {
	create  *CreateRepository
	update  *UpdateRepository
	delete  *DeleteRepository
	getById *GetByIdRepository
	getList *GetListRepository
}

func NewRepository(config *entity.Config, redisClient *redis.Client, postgresClient *services.PostgresClient) contacts.Repository {
	return &contactsRepo{
		create:  NewCreateRepository(config, redisClient, postgresClient),
		update:  NewUpdateRepository(config, redisClient, postgresClient),
		delete:  NewDeleteRepository(config, redisClient, postgresClient),
		getById: NewGetByIdRepository(config, redisClient, postgresClient),
		getList: NewGetListRepository(config, redisClient, postgresClient),
	}
}

func (r *contactsRepo) Create(ctx context.Context, req entity.Contact) (*int, error) {
	return r.create.Execute(ctx, req)
}

func (r *contactsRepo) Update(ctx context.Context, req entity.Contact) (*int, error) {
	return r.update.Execute(ctx, req)
}

func (r *contactsRepo) Delete(ctx context.Context, id int) error {
	return r.delete.Execute(ctx, id)
}

func (r *contactsRepo) GetById(ctx context.Context, id int) (*entity.Contact, error) {
	return r.getById.Execute(ctx, id)
}

func (r *contactsRepo) GetList(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.Contact], error) {
	return r.getList.Execute(ctx, listReq)
}
