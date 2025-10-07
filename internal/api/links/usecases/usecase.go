package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/links"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type linksUC struct {
	upsert  *UpsertUC
	getListById *GetListByIdUC
	deleteAll *DeleteAllUC
}

func NewUseCase(config *entity.Config, linksRepo links.Repository, postgresClient *services.PostgresClient) links.UseCase {
	return &linksUC{
		upsert:  NewUpsertUC(config, linksRepo, postgresClient),
		getListById: NewGetListByIdUC(config, linksRepo, postgresClient),
		deleteAll: NewDeleteAllUC(config, linksRepo, postgresClient),
	}
}

func (u *linksUC) Upsert(ctx context.Context, input entity.LinkArray) error {
	return u.upsert.Execute(ctx, input)
}

func (u *linksUC) GetListById(ctx context.Context, module string, parent_id int) (entity.LinkRespArray, error) {
	return u.getListById.Execute(ctx, module, parent_id)
}

func (u *linksUC) DeleteAll(ctx context.Context, module string, parent_id int) error {
	return u.deleteAll.Execute(ctx, module, parent_id)
}