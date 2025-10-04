package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jip/portfolio-backend/internal/api/experiences"
	"github.com/jip/portfolio-backend/internal/entity"
)

type experiencesUC struct {
	create 	*CreateUC
	update 	*UpdateUC
	delete 	*DeleteUC
	getById *GetByIdUC
	getList *GetListUC
}

func NewExperiencesUseCase(config *entity.Config, redisClient *redis.Client, experiencesRepo experiences.Repository) experiences.UseCase {
	byId := NewGetByIdUC(config, redisClient, experiencesRepo)
	
	return &experiencesUC{
		create: 	NewCreateUC(config, redisClient, experiencesRepo, byId),
		update: 	NewUpdateUC(config, redisClient, experiencesRepo, byId),
		delete: 	NewDeleteUC(config, redisClient, experiencesRepo),
		getById: 	byId,
		getList: 	NewGetListUC(config, redisClient, experiencesRepo),
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

func (u *experiencesUC) GetList(ctx context.Context, ListReq entity.ListReq) (*entity.List, error) {
	return u.getList.Execute(ctx, ListReq)
}