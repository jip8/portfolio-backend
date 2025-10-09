package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/projects"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/jip/portfolio-backend/internal/api/links"
)

var (
	ModuleName = "projects"
)

type projectsUC struct {
	create  *CreateUC
	update  *UpdateUC
	delete  *DeleteUC
	getById *GetByIdUC
	getList *GetListUC
}

func NewUseCase(config *entity.Config, projectsRepo projects.Repository, postgresClient *services.PostgresClient, linksUC links.UseCase) projects.UseCase {
	byId := NewGetByIdUC(config, projectsRepo, postgresClient, linksUC)

	return &projectsUC{
		create:  NewCreateUC(config, projectsRepo, byId, postgresClient, linksUC),
		update:  NewUpdateUC(config, projectsRepo, byId, postgresClient, linksUC),
		delete:  NewDeleteUC(config, projectsRepo, postgresClient, linksUC),
		getById: byId,
		getList: NewGetListUC(config, projectsRepo, postgresClient),
	}
}

func (u *projectsUC) Create(ctx context.Context, req entity.ProjectFlat) (*entity.ProjectResp, error) {
	return u.create.Execute(ctx, req)
}

func (u *projectsUC) Update(ctx context.Context, req entity.ProjectFlat) (*entity.ProjectResp, error) {
	return u.update.Execute(ctx, req)
}

func (u *projectsUC) Delete(ctx context.Context, id int) error {
	return u.delete.Execute(ctx, id)
}

func (u *projectsUC) GetById(ctx context.Context, id int) (*entity.ProjectResp, error) {
	return u.getById.Execute(ctx, id)
}

func (u *projectsUC) GetList(ctx context.Context, ListReq entity.ListReq) (*entity.List[entity.ProjectResp], error) {
	return u.getList.Execute(ctx, ListReq)
}
