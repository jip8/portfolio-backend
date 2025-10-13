package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/attachments"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type attachmentsUC struct {
	insert      *InsertUC
	getListById *GetListByIdUC
	deleteAll   *DeleteAllUC
	delete      *DeleteUC
	get         *GetUC
	minioClient *services.MinioClient
}

func NewUseCase(config *entity.Config, attachmentsRepo attachments.Repository, postgresClient *services.PostgresClient, minioClient *services.MinioClient) attachments.UseCase {
	return &attachmentsUC{
		insert:      NewInsertUC(config, attachmentsRepo, postgresClient, minioClient),
		getListById: NewGetListByIdUC(config, attachmentsRepo, postgresClient),
		deleteAll:   NewDeleteAllUC(config, attachmentsRepo, postgresClient, minioClient),
		delete:      NewDeleteUC(config, attachmentsRepo, postgresClient, minioClient),
		get:         NewGetUC(config, attachmentsRepo, postgresClient, minioClient),
		minioClient: minioClient,
	}
}

func (u *attachmentsUC) Insert(ctx context.Context, input []entity.AttachmentFlat) error {
	return u.insert.Execute(ctx, input)
}

func (u *attachmentsUC) GetListById(ctx context.Context, module string, parent_id int) (entity.AttachmentRespArray, error) {
	return u.getListById.Execute(ctx, module, parent_id)
}

func (u *attachmentsUC) DeleteAll(ctx context.Context, module string, parent_id int) error {
	return u.deleteAll.Execute(ctx, module, parent_id)
}

func (u *attachmentsUC) Delete(ctx context.Context, module string, parent_id int, ids []int) error {
	return u.delete.Execute(ctx, module, parent_id, ids)
}

func (u *attachmentsUC) Get(ctx context.Context, objectKey string) (*entity.File, error) {
	return u.get.Execute(ctx, objectKey)
}