package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/entity"
)

type GetListUC struct {
	config      *entity.Config
	redisClient *redis.Client
}

func NewGetListUC(config *entity.Config, redisClient *redis.Client) *GetListUC {
	return &GetListUC{
		config:      config,
		redisClient: redisClient,
	}
}

func (u *GetListUC) Execute(ctx context.Context, ListReq entity.ListReq) (*entity.List, error) {
	

	return nil, nil
}