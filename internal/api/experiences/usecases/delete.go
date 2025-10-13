package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/experiences"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/skills"
)

type DeleteUC struct {
	config          *entity.Config
	experiencesRepo experiences.Repository
	postgresClient 	*services.PostgresClient
	skillsUC       	skills.UseCase
}

func NewDeleteUC(config *entity.Config, experiencesRepo experiences.Repository, postgresClient *services.PostgresClient, skillsUC skills.UseCase) *DeleteUC {
	return &DeleteUC{
		config:          	config,
		experiencesRepo: 	experiencesRepo,
		postgresClient: 	postgresClient,
		skillsUC:       	skillsUC,
	}
}

func (u *DeleteUC) Execute(ctx context.Context, id int) (err error) {
	ctx, err = u.postgresClient.StartProcess(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = u.postgresClient.CloseProcess(ctx, err)
	}()

	err = u.experiencesRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	err = u.skillsUC.DeleteAll(ctx, moduleName, id)
	if err != nil {
		return err
	}

	return nil
}