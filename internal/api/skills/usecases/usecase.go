package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/skills"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type skillsUC struct {
	upsert     		*UpsertUC
	getListById 	*GetListByIdUC
	getList 		*GetListUC
	deleteAll   	*DeleteAllUC
}

func NewUseCase(config *entity.Config, skillsRepo skills.Repository, postgresClient *services.PostgresClient) skills.UseCase {
	return &skillsUC{
		upsert:     	NewUpsertUC(config, skillsRepo, postgresClient),
		getList: 		NewGetListUC(config, skillsRepo, postgresClient),
		getListById: 	NewGetListByIdUC(config, skillsRepo, postgresClient),
		deleteAll:  	NewDeleteAllUC(config, skillsRepo, postgresClient),
	}
}

func (u *skillsUC) Upsert(ctx context.Context, parent_id *int, module *string, input entity.SkillArray) error {
	return u.upsert.Execute(ctx, parent_id, module, input)
}

func (u *skillsUC) GetListById(ctx context.Context, module *string, parent_id *int) (entity.SkillRespArray, error) {
	return u.getListById.Execute(ctx, module, parent_id)
}

func (u *skillsUC) GetList(ctx context.Context) (entity.SkillRespArray, error) {
	return u.getList.Execute(ctx)
}

func (u *skillsUC) DeleteAll(ctx context.Context, module string, parent_id int) error {
	return u.deleteAll.Execute(ctx, module, parent_id)
}
