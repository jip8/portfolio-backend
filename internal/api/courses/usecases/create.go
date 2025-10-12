package usecases

import (
	"context"
	"fmt"

	"github.com/jip/portfolio-backend/internal/api/courses"
	"github.com/jip/portfolio-backend/internal/api/skills"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type CreateUC struct {
	config         *entity.Config
	coursesRepo    courses.Repository
	byId           *GetByIdUC
	postgresClient *services.PostgresClient
	skillsUC       skills.UseCase
}

func NewCreateUC(config *entity.Config, coursesRepo courses.Repository, byId *GetByIdUC, postgresClient *services.PostgresClient, skillsUC skills.UseCase) *CreateUC {
	return &CreateUC{
		config:         config,
		coursesRepo:    coursesRepo,
		byId:           byId,
		postgresClient: postgresClient,
		skillsUC:       skillsUC,
	}
}

func (u *CreateUC) Execute(ctx context.Context, req entity.CourseFlat) (resp *entity.CourseResp, err error) {
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
	createdId, err = u.coursesRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	module := fmt.Sprintf("%s", moduleName)
	err = u.skillsUC.Upsert(ctx, createdId, &module, req.Skills)
	if err != nil {
		return nil, err
	}

	resp, err = u.byId.Execute(ctx, *createdId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
