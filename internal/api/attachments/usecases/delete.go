package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/attachments"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
	"github.com/minio/minio-go/v7"
)

type DeleteUC struct {
	config          *entity.Config
	attachmentsRepo attachments.Repository
	postgresClient  *services.PostgresClient
	minioClient     *services.MinioClient
}

func NewDeleteUC(config *entity.Config, attachmentsRepo attachments.Repository, postgresClient *services.PostgresClient, minioClient *services.MinioClient) *DeleteUC {
	return &DeleteUC{
		config:          config,
		attachmentsRepo: attachmentsRepo,
		postgresClient:  postgresClient,
		minioClient:     minioClient,
	}
}

func (u *DeleteUC) Execute(ctx context.Context, module string, parent_id int, ids []int) error {
	var err error

	allAttachments, err := u.attachmentsRepo.GetListById(ctx, module, parent_id)
	if err != nil {
		return err
	}

	attachmentsToDelete := []entity.AttachmentResp{}
	for _, att := range allAttachments {
		for _, id := range ids {
			if att.Id == id {
				attachmentsToDelete = append(attachmentsToDelete, att)
				break
			}
		}
	}

	for _, att := range attachmentsToDelete {
		_ = u.minioClient.RemoveObject(ctx, att.Link, minio.RemoveObjectOptions{})
	}

	ctx, err = u.postgresClient.StartProcess(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = u.postgresClient.CloseProcess(ctx, err)
	}()

	err = u.attachmentsRepo.Delete(ctx, module, parent_id, ids)
	if err != nil {
		return err
	}

	return nil
}