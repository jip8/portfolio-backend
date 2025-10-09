package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/experiences"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type experiencesRepo struct {
	create  *CreateRepository
	update  *UpdateRepository
	delete  *DeleteRepository
	getById *GetByIdRepository
	getList *GetListRepository
}

func NewRepository(config *entity.Config, postgresClient *services.PostgresClient) experiences.Repository {
	return &experiencesRepo{
		create:  NewCreateRepository(config, postgresClient),
		update:  NewUpdateRepository(config, postgresClient),
		delete:  NewDeleteRepository(config, postgresClient),
		getById: NewGetByIdRepository(config, postgresClient),
		getList: NewGetListRepository(config, postgresClient),
	}
}

func (r *experiencesRepo) Create(ctx context.Context, req entity.ExperienceFlat) (*int, error) {
	return r.create.Execute(ctx, req)
}

func (r *experiencesRepo) Update(ctx context.Context, req entity.ExperienceFlat) (*int, error) {
	return r.update.Execute(ctx, req)
}

func (r *experiencesRepo) Delete(ctx context.Context, id int) error {
	return r.delete.Execute(ctx, id)
}

func (r *experiencesRepo) GetById(ctx context.Context, id int) (*entity.ExperienceResp, error) {
	return r.getById.Execute(ctx, id)
}

func (r *experiencesRepo) GetList(ctx context.Context, listReq entity.ListReq) (*entity.List[entity.ExperienceResp], error) {
	return r.getList.Execute(ctx, listReq)
}