package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/experiences"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/skills"
)

const (
	moduleName = "experiences"
)

type experiencesUC struct {
	create  *CreateUC
	update  *UpdateUC
	delete  *DeleteUC
	getById *GetByIdUC
	getList *GetListUC
}


func NewUseCase(config *entity.Config, experiencesRepo experiences.Repository, postgresClient *services.PostgresClient, skillsUC skills.UseCase) experiences.UseCase {
	byId := NewGetByIdUC(config, experiencesRepo, postgresClient, skillsUC)

	return &experiencesUC{
		create:  NewCreateUC(config, experiencesRepo, byId, postgresClient, skillsUC),
		update:  NewUpdateUC(config, experiencesRepo, byId, postgresClient, skillsUC),
		delete:  NewDeleteUC(config, experiencesRepo, postgresClient, skillsUC),
		getById: byId,
		getList: NewGetListUC(config, experiencesRepo, postgresClient),
	}
}

func (u *experiencesUC) Create(ctx context.Context, req entity.ExperienceFlat) (*entity.ExperienceResp, error) {
	return u.create.Execute(ctx, req)
}

func (u *experiencesUC) Update(ctx context.Context, req entity.ExperienceFlat) (*entity.ExperienceResp, error) {
	return u.update.Execute(ctx, req)
}

func (u *experiencesUC) Delete(ctx context.Context, id int) error {
	return u.delete.Execute(ctx, id)
}

func (u *experiencesUC) GetById(ctx context.Context, id int) (*entity.ExperienceResp, error) {
	return u.getById.Execute(ctx, id)
}

func (u *experiencesUC) GetList(ctx context.Context, ListReq entity.ListReq) (*entity.List[entity.ExperienceResp], error) {
	return u.getList.Execute(ctx, ListReq)
}