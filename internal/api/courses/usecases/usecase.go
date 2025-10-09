package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/courses"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type coursesUC struct {
	create  *CreateUC
	update  *UpdateUC
	delete  *DeleteUC
	getById *GetByIdUC
	getList *GetListUC
}

func NewUseCase(config *entity.Config, coursesRepo courses.Repository, postgresClient *services.PostgresClient) courses.UseCase {
	byId := NewGetByIdUC(config, coursesRepo, postgresClient)

	return &coursesUC{
		create:  NewCreateUC(config, coursesRepo, byId, postgresClient),
		update:  NewUpdateUC(config, coursesRepo, byId, postgresClient),
		delete:  NewDeleteUC(config, coursesRepo, postgresClient),
		getById: byId,
		getList: NewGetListUC(config, coursesRepo, postgresClient),
	}
}

func (u *coursesUC) Create(ctx context.Context, req entity.CourseFlat) (*entity.CourseResp, error) {
	return u.create.Execute(ctx, req)
}

func (u *coursesUC) Update(ctx context.Context, req entity.CourseFlat) (*entity.CourseResp, error) {
	return u.update.Execute(ctx, req)
}

func (u *coursesUC) Delete(ctx context.Context, id int) error {
	return u.delete.Execute(ctx, id)
}

func (u *coursesUC) GetById(ctx context.Context, id int) (*entity.CourseResp, error) {
	return u.getById.Execute(ctx, id)
}

func (u *coursesUC) GetList(ctx context.Context, ListReq entity.ListReq) (*entity.List[entity.CourseResp], error) {
	return u.getList.Execute(ctx, ListReq)
}
