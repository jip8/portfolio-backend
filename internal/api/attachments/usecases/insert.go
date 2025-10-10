package usecases

import (
	"context"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/jip/portfolio-backend/internal/api/attachments"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type InsertUC struct {
	config          *entity.Config
	attachmentsRepo attachments.Repository
	postgresClient  *services.PostgresClient
	minioClient     *services.MinioClient
}

func NewInsertUC(config *entity.Config, attachmentsRepo attachments.Repository, postgresClient *services.PostgresClient, minioClient *services.MinioClient) *InsertUC {
	return &InsertUC{
		config:          config,
		attachmentsRepo: attachmentsRepo,
		postgresClient:  postgresClient,
		minioClient:     minioClient,
	}
}

func (u *InsertUC) Execute(ctx context.Context, input []entity.AttachmentFlat) error {
	var err error

	for i := range input {
		if input[i].FileObject == nil {
			continue
		}
		file := input[i].FileObject
		fileName := uuid.New().String() + filepath.Ext(file.Name)
		file.Name = fileName

		uploadInfo, err := u.minioClient.AddObject(ctx, file)
		if err != nil {
			return err
		}
		input[i].Link = uploadInfo.Key
		input[i].ContentType = &file.ContentType
	}

	ctx, err = u.postgresClient.StartProcess(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = u.postgresClient.CloseProcess(ctx, err)
	}()

	err = u.attachmentsRepo.Insert(ctx, input)
	if err != nil {
		return err
	}

	return nil
}