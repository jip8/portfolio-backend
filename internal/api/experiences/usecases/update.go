package usecases

import (
	"fmt"
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/api/experiences"

	"github.com/jip/portfolio-backend"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/skills"
)

type UpdateUC struct {
	config      	*entity.Config
	experiencesRepo experiences.Repository
	byId 			*GetByIdUC
	postgresClient 	*services.PostgresClient
	skillsUC      	skills.UseCase
}

func NewUpdateUC(config *entity.Config, experiencesRepo experiences.Repository, byId *GetByIdUC, postgresClient *services.PostgresClient, skillsUC skills.UseCase) *UpdateUC {
	return &UpdateUC{
		config:      		config,
		experiencesRepo: 	experiencesRepo,
		byId: 				byId,
		postgresClient: 	postgresClient,
		skillsUC:      		skillsUC,
	}
}

func (u *UpdateUC) Execute(ctx context.Context, req entity.ExperienceFlat) (resp *entity.ExperienceResp, err error) {
	if req.Id == nil {
		return nil, portfolio.ErrExperienceIdIsRequired
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
	updatedId, err = u.experiencesRepo.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	module := fmt.Sprintf("%s", moduleName)
	err = u.skillsUC.Upsert(ctx, updatedId, &module, req.Skills)
	if err != nil {
		return nil, err
	}

	resp, err = u.byId.Execute(ctx, *updatedId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}