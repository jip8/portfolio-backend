package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/courses"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type coursesRepo struct {
	create  *CreateRepository
	update  *UpdateRepository
	delete  *DeleteRepository
	getById *GetByIdRepository
	getList *GetListRepository
}

func NewRepository(config *entity.Config, postgresClient *services.PostgresClient) courses.Repository {
	return &coursesRepo{
		create:  NewCreateRepository(config, postgresClient),
		update:  NewUpdateRepository(config, postgresClient),
		delete:  NewDeleteRepository(config, postgresClient),
		getById: NewGetByIdRepository(config, postgresClient),
		getList: NewGetListRepository(config, postgresClient),
	}
}

func (r *coursesRepo) Create(ctx context.Context, req entity.CourseFlat) (*int, error) {
	return r.create.Execute(ctx, req)
}

func (r *coursesRepo) Update(ctx context.Context, req entity.CourseFlat) (*int, error) {
	return r.update.Execute(ctx, req)
}

func (r *coursesRepo) Delete(ctx context.Context, id int) error {
	return r.delete.Execute(ctx, id)
}

func (r *coursesRepo) GetById(ctx context.Context, id int) (*entity.CourseResp, error) {
	return r.getById.Execute(ctx, id)
}

func (r *coursesRepo) GetList(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.CourseResp], error) {
	return r.getList.Execute(ctx, listReq)
}
