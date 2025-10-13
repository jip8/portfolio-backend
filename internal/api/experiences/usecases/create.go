package usecases

import (
	"fmt"
	"context"

	"github.com/jip/portfolio-backend/internal/api/experiences"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/skills"
)

type CreateUC struct {
	config      	*entity.Config
	experiencesRepo experiences.Repository
	byId 			*GetByIdUC
	postgresClient 	*services.PostgresClient
	skillsUC      	skills.UseCase
}

func NewCreateUC(config *entity.Config, experiencesRepo experiences.Repository, byId *GetByIdUC, postgresClient *services.PostgresClient, skillsUC skills.UseCase) *CreateUC {
	return &CreateUC{
		config:      		config,
		experiencesRepo: 	experiencesRepo,
		byId: 				byId,
		postgresClient: 	postgresClient,
		skillsUC:      	skillsUC,
	}
}

func (u *CreateUC) Execute(ctx context.Context, req entity.ExperienceFlat) (resp *entity.ExperienceResp, err error) {
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
	createdId, err = u.experiencesRepo.Create(ctx, req)
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