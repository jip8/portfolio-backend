package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/links"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type linksRepo struct {
	upsert      *UpsertRepository
	delete      *DeleteRepository
	deleteAll   *DeleteAllRepository
	getListById *GetListByIdRepository
}

func NewRepository(config *entity.Config, postgresClient *services.PostgresClient) links.Repository {
	return &linksRepo{
		upsert:      NewUpsertRepository(config, postgresClient),
		delete:      NewDeleteRepository(config, postgresClient),
		deleteAll:   NewDeleteAllRepository(config, postgresClient),
		getListById: NewGetListByIdRepository(config, postgresClient),
	}
}

func (r *linksRepo) Upsert(ctx context.Context, input entity.LinkArray) error {
	return r.upsert.Execute(ctx, input)
}

func (r *linksRepo) Delete(ctx context.Context, ids []int) error {
	return r.delete.Execute(ctx, ids)
}

func (r *linksRepo) DeleteAll(ctx context.Context, module string, parent_id int) error {
	return r.deleteAll.Execute(ctx, module, parent_id)
}

func (r *linksRepo) GetListById(ctx context.Context, module string, parent_id int) (entity.LinkRespArray, error) {
	return r.getListById.Execute(ctx, module, parent_id)
}
