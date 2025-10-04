package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"github.com/jip/portfolio-backend/internal/api/experiences"
)

type experiencesRepo struct {
	create 	*CreateRepository
	update 	*UpdateRepository
	delete 	*DeleteRepository
	getById *GetByIdRepository
	getList *GetListRepository
}

func NewExperiencesRepository(config *entity.Config, redisClient *redis.Client, db *gorm.DB) experiences.Repository {
	return &experiencesRepo{
		create: NewCreateRepository(config, redisClient, db),
		update: NewUpdateRepository(config, redisClient, db),
		delete: NewDeleteRepository(config, redisClient, db),
		getById: NewGetByIdRepository(config, redisClient, db),
		getList: NewGetListRepository(config, redisClient, db),
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

func (r *experiencesRepo) GetList(ctx context.Context, listReq entity.ListReq) (*entity.List, error) {
	return r.getList.Execute(ctx, listReq)
}