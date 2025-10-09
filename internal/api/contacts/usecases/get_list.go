package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/contacts"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetListUC struct {
	config         *entity.Config
	redisClient    *redis.Client
	contactsRepo   contacts.Repository
	postgresClient *services.PostgresClient
}

func NewGetListUC(config *entity.Config,  contactsRepo contacts.Repository, postgresClient *services.PostgresClient) *GetListUC {
	return &GetListUC{
		config:         config,
		contactsRepo:   contactsRepo,
		postgresClient: postgresClient,
	}
}

func (u *GetListUC) Execute(ctx context.Context, ListReq entity.ListReq) (*entity.List[entity.Contact], error) {

	resp, err := u.contactsRepo.GetList(ctx, ListReq)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
