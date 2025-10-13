package repositories

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/attachments"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type attachmentsRepo struct {
	insert      *InsertRepository
	delete      *DeleteRepository
	deleteAll   *DeleteAllRepository
	getListById *GetListByIdRepository
}

func NewRepository(config *entity.Config, postgresClient *services.PostgresClient) attachments.Repository {
	return &attachmentsRepo{
		insert:      NewInsertRepository(config, postgresClient),
		delete:      NewDeleteRepository(config, postgresClient),
		deleteAll:   NewDeleteAllRepository(config, postgresClient),
		getListById: NewGetListByIdRepository(config, postgresClient),
	}
}

func (r *attachmentsRepo) Insert(ctx context.Context, input []entity.AttachmentFlat) error {
	return r.insert.Execute(ctx, input)
}

func (r *attachmentsRepo) Delete(ctx context.Context, module string, parent_id int, ids []int) error {
	return r.delete.Execute(ctx, module, parent_id, ids)
}

func (r *attachmentsRepo) DeleteAll(ctx context.Context, module string, parent_id int) error {
	return r.deleteAll.Execute(ctx, module, parent_id)
}

func (r *attachmentsRepo) GetListById(ctx context.Context, module string, parent_id int) (entity.AttachmentRespArray, error) {
	return r.getListById.Execute(ctx, module, parent_id)
}