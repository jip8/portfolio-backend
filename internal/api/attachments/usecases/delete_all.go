package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/attachments"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/minio/minio-go/v7"
)

type DeleteAllUC struct {
	config          *entity.Config
	attachmentsRepo attachments.Repository
	postgresClient  *services.PostgresClient
	minioClient     *services.MinioClient
}

func NewDeleteAllUC(config *entity.Config, attachmentsRepo attachments.Repository, postgresClient *services.PostgresClient, minioClient *services.MinioClient) *DeleteAllUC {
	return &DeleteAllUC{
		config:          config,
		attachmentsRepo: attachmentsRepo,
		postgresClient:  postgresClient,
		minioClient:     minioClient,
	}
}

func (u *DeleteAllUC) Execute(ctx context.Context, module string, parent_id int) error {
	attachments, err := u.attachmentsRepo.GetListById(ctx, module, parent_id)
	if err != nil {
		return err
	}

	for _, att := range attachments {
		_ = u.minioClient.RemoveObject(ctx, att.Link, minio.RemoveObjectOptions{})
	}

	err = u.attachmentsRepo.DeleteAll(ctx, module, parent_id)
	if err != nil {
		return err
	}

	return nil
}