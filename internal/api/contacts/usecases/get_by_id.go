package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/contacts"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetByIdUC struct {
	config         *entity.Config
	redisClient    *redis.Client
	contactsRepo   contacts.Repository
	postgresClient *services.PostgresClient
}

func NewGetByIdUC(config *entity.Config, redisClient *redis.Client, contactsRepo contacts.Repository, postgresClient *services.PostgresClient) *GetByIdUC {
	return &GetByIdUC{
		config:         config,
		redisClient:    redisClient,
		contactsRepo:   contactsRepo,
		postgresClient: postgresClient,
	}
}

func (u *GetByIdUC) Execute(ctx context.Context, id int) (*entity.Contact, error) {

	resp, err := u.contactsRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
