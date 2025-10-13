package usecases

import (
	"context"

	"github.com/jip/portfolio-backend/internal/api/attachments"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/jip/portfolio-backend/internal/services"
)

type GetUC struct {
	config          *entity.Config
	attachmentsRepo attachments.Repository
	postgresClient  *services.PostgresClient
	minioClient     *services.MinioClient
}

func NewGetUC(config *entity.Config, attachmentsRepo attachments.Repository, postgresClient *services.PostgresClient, minioClient *services.MinioClient) *GetUC {
	return &GetUC{
		config:          config,
		attachmentsRepo: attachmentsRepo,
		postgresClient:  postgresClient,
		minioClient:     minioClient,
	}
}

func (u *GetUC) Execute(ctx context.Context, objectKey string) (*entity.File, error) {
	obj, err := u.minioClient.GetObject(ctx, objectKey)
	if err != nil {
		return nil, err
	}

	stat, err := obj.Stat()
	if err != nil {
		return nil, err
	}

	return &entity.File{
		Name:        stat.Key,
		Size:        stat.Size,
		ContentType: stat.ContentType,
		Content:     obj,
	}, nil
}