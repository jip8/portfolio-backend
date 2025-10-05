package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/courses"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetListUC struct {
	config         *entity.Config
	redisClient    *redis.Client
	coursesRepo    courses.Repository
	postgresClient *services.PostgresClient
}

func NewGetListUC(config *entity.Config, redisClient *redis.Client, coursesRepo courses.Repository, postgresClient *services.PostgresClient) *GetListUC {
	return &GetListUC{
		config:         config,
		redisClient:    redisClient,
		coursesRepo:    coursesRepo,
		postgresClient: postgresClient,
	}
}

func (u *GetListUC) Execute(ctx context.Context, ListReq entity.ListReq) (*entity.List[entity.CourseResp], error) {

	resp, err := u.coursesRepo.GetList(ctx, ListReq)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
