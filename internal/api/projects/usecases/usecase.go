package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/projects"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type projectsUC struct {
	create  *CreateUC
	update  *UpdateUC
	delete  *DeleteUC
	getById *GetByIdUC
	getList *GetListUC
}

func NewUseCase(config *entity.Config, redisClient *redis.Client, projectsRepo projects.Repository, postgresClient *services.PostgresClient) projects.UseCase {
	byId := NewGetByIdUC(config, redisClient, projectsRepo, postgresClient)

	return &projectsUC{
		create:  NewCreateUC(config, redisClient, projectsRepo, byId, postgresClient),
		update:  NewUpdateUC(config, redisClient, projectsRepo, byId, postgresClient),
		delete:  NewDeleteUC(config, redisClient, projectsRepo, postgresClient),
		getById: byId,
		getList: NewGetListUC(config, redisClient, projectsRepo, postgresClient),
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
