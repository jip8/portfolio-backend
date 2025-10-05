package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/contacts"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type CreateUC struct {
	config         *entity.Config
	redisClient    *redis.Client
	contactsRepo   contacts.Repository
	byId           *GetByIdUC
	postgresClient *services.PostgresClient
}

func NewCreateUC(config *entity.Config, redisClient *redis.Client, contactsRepo contacts.Repository, byId *GetByIdUC, postgresClient *services.PostgresClient) *CreateUC {
	return &CreateUC{
		config:         config,
		redisClient:    redisClient,
		contactsRepo:   contactsRepo,
		byId:           byId,
		postgresClient: postgresClient,
	}
}

func (u *CreateUC) Execute(ctx context.Context, req entity.Contact) (resp *entity.Contact, err error) {
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

	var createdId *int
	createdId, err = u.contactsRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	resp, err = u.byId.Execute(ctx, *createdId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
