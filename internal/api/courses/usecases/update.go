package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/courses"
	"github.com/jip/portfolio-backend/internal/entity"

	"github.com/jip/portfolio-backend"
	"github.com/jip/portfolio-backend/internal/services"
)

type UpdateUC struct {
	config         *entity.Config
	redisClient    *redis.Client
	coursesRepo    courses.Repository
	byId           *GetByIdUC
	postgresClient *services.PostgresClient
}

func NewUpdateUC(config *entity.Config, redisClient *redis.Client, coursesRepo courses.Repository, byId *GetByIdUC, postgresClient *services.PostgresClient) *UpdateUC {
	return &UpdateUC{
		config:         config,
		redisClient:    redisClient,
		coursesRepo:    coursesRepo,
		byId:           byId,
		postgresClient: postgresClient,
	}
}

func (u *UpdateUC) Execute(ctx context.Context, req entity.CourseFlat) (resp *entity.CourseResp, err error) {
	if req.Id == nil {
		return nil, portfolio.ErrCourseIdIsRequired
	}

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

	var updatedId *int
	updatedId, err = u.coursesRepo.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	resp, err = u.byId.Execute(ctx, *updatedId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
