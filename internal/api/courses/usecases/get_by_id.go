package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/courses"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetByIdUC struct {
	config         *entity.Config
	coursesRepo    courses.Repository
	postgresClient *services.PostgresClient
}

func NewGetByIdUC(config *entity.Config, coursesRepo courses.Repository, postgresClient *services.PostgresClient) *GetByIdUC {
	return &GetByIdUC{
		config:         config,
		coursesRepo:    coursesRepo,
		postgresClient: postgresClient,
	}
}

func (u *GetByIdUC) Execute(ctx context.Context, id int) (*entity.CourseResp, error) {

	resp, err := u.coursesRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		err = resp.Format()
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}
