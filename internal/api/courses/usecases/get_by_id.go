package usecases

import (
	"context"
	"fmt"

	"github.com/jip/portfolio-backend/internal/api/courses"
	"github.com/jip/portfolio-backend/internal/api/skills"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetByIdUC struct {
	config         *entity.Config
	coursesRepo    courses.Repository
	postgresClient *services.PostgresClient
	skillsUC       skills.UseCase
}

func NewGetByIdUC(config *entity.Config, coursesRepo courses.Repository, postgresClient *services.PostgresClient, skillsUC skills.UseCase) *GetByIdUC {
	return &GetByIdUC{
		config:         config,
		coursesRepo:    coursesRepo,
		postgresClient: postgresClient,
		skillsUC:       skillsUC,
	}
}

func (u *GetByIdUC) Execute(ctx context.Context, id int) (*entity.CourseResp, error) {

	resp, err := u.coursesRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	module := fmt.Sprintf("%s", moduleName)
	skills, err := u.skillsUC.GetList(ctx, &module, &id)
	if err != nil {
		return nil, err
	}

	resp.Skills = skills

	if resp != nil {
		err = resp.Format()
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}
