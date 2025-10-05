package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/contacts"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type contactsUC struct {
	create  *CreateUC
	update  *UpdateUC
	delete  *DeleteUC
	getById *GetByIdUC
	getList *GetListUC
}

func NewUseCase(config *entity.Config, redisClient *redis.Client, contactsRepo contacts.Repository, postgresClient *services.PostgresClient) contacts.UseCase {
	byId := NewGetByIdUC(config, redisClient, contactsRepo, postgresClient)

	return &contactsUC{
		create:  NewCreateUC(config, redisClient, contactsRepo, byId, postgresClient),
		update:  NewUpdateUC(config, redisClient, contactsRepo, byId, postgresClient),
		delete:  NewDeleteUC(config, redisClient, contactsRepo, postgresClient),
		getById: byId,
		getList: NewGetListUC(config, redisClient, contactsRepo, postgresClient),
	}
}

func (u *contactsUC) Create(ctx context.Context, req entity.Contact) (*entity.Contact, error) {
	return u.create.Execute(ctx, req)
}

func (u *contactsUC) Update(ctx context.Context, req entity.Contact) (*entity.Contact, error) {
	return u.update.Execute(ctx, req)
}

func (u *contactsUC) Delete(ctx context.Context, id int) error {
	return u.delete.Execute(ctx, id)
}

func (u *contactsUC) GetById(ctx context.Context, id int) (*entity.Contact, error) {
	return u.getById.Execute(ctx, id)
}

func (u *contactsUC) GetList(ctx context.Context, ListReq entity.ListReq) (*entity.List[entity.Contact], error) {
	return u.getList.Execute(ctx, ListReq)
}
