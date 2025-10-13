package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/skills"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type skillsRepo struct {
	upsert      	*UpsertRepository
	delete      	*DeleteRepository
	deleteAll   	*DeleteAllRepository
	getListById 	*GetListByIdRepository
	getList     	*GetListRepository
	addExclusive 	*AddExclusiveRepository
}

func NewRepository(config *entity.Config, postgresClient *services.PostgresClient) skills.Repository {
	return &skillsRepo{
		upsert:      	NewUpsertRepository(config, postgresClient),
		delete:      	NewDeleteRepository(config, postgresClient),
		deleteAll:   	NewDeleteAllRepository(config, postgresClient),
		getListById: 	NewGetListByIdRepository(config, postgresClient),
		getList:     	NewGetListRepository(config, postgresClient),
		addExclusive: 	NewAddExclusiveRepository(config, postgresClient),
	}
}

func (r *skillsRepo) Upsert(ctx context.Context, input entity.SkillArray) ([]int, error) {
	return r.upsert.Execute(ctx, input)
}

func (r *skillsRepo) Delete(ctx context.Context, ids []int) error {
	return r.delete.Execute(ctx, ids)
}

func (r *skillsRepo) DeleteAll(ctx context.Context, module string, parent_id int) error {
	return r.deleteAll.Execute(ctx, module, parent_id)
}

func (r *skillsRepo) GetListById(ctx context.Context, module *string, parent_id *int) (entity.SkillRespArray, error) {
	return r.getListById.Execute(ctx, module, parent_id)
}

func (r *skillsRepo) GetList(ctx context.Context) (entity.SkillRespArray, error) {
	return r.getList.Execute(ctx)
}

func (r *skillsRepo) AddExclusive(ctx context.Context, parent_id *int, module *string, ids []int) error {
	return r.addExclusive.Execute(ctx, parent_id, module, ids)
}
