package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/projects"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type projectsRepo struct {
	create  *CreateRepository
	update  *UpdateRepository
	delete  *DeleteRepository
	getById *GetByIdRepository
	getList *GetListRepository
}

func NewRepository(config *entity.Config, postgresClient *services.PostgresClient) projects.Repository {
	return &projectsRepo{
		create:  NewCreateRepository(config, postgresClient),
		update:  NewUpdateRepository(config, postgresClient),
		delete:  NewDeleteRepository(config, postgresClient),
		getById: NewGetByIdRepository(config, postgresClient),
		getList: NewGetListRepository(config, postgresClient),
	}
}

func (r *projectsRepo) Create(ctx context.Context, req entity.ProjectFlat) (*int, error) {
	return r.create.Execute(ctx, req)
}

func (r *projectsRepo) Update(ctx context.Context, req entity.ProjectFlat) (*int, error) {
	return r.update.Execute(ctx, req)
}

func (r *projectsRepo) Delete(ctx context.Context, id int) error {
	return r.delete.Execute(ctx, id)
}

func (r *projectsRepo) GetById(ctx context.Context, id int) (*entity.ProjectResp, error) {
	return r.getById.Execute(ctx, id)
}

func (r *projectsRepo) GetList(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.ProjectResp], error) {
	return r.getList.Execute(ctx, listReq)
}
