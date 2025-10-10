package services

import (
	"context"
	"path"

	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client *minio.Client
	Bucket string
}

func NewMinioClient(cfg *entity.Config) (*MinioClient, error) {
	endpoint := cfg.Minio.Endpoint
	accessKeyID := cfg.Minio.AccessKeyID
	secretAccessKey := cfg.Minio.SecretAccessKey
	useSSL := cfg.Minio.UseSSL
	bucket := cfg.Minio.Bucket

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &MinioClient{
		Client: client,
		Bucket: bucket,
	}, nil
}

func (m *MinioClient) AddObject(ctx context.Context, file *entity.File) (minio.UploadInfo, error) {
	opts := minio.PutObjectOptions{
		ContentType: file.ContentType,
	}
	objectName := path.Join("attachments", file.Name)
	return m.Client.PutObject(ctx, m.Bucket, objectName, file.Content, file.Size, opts)
}

func (m *MinioClient) GetObject(ctx context.Context, objectName string) (*minio.Object, error) {
	fullObjectName := path.Join("attachments", objectName)
	return m.Client.GetObject(ctx, m.Bucket, fullObjectName, minio.GetObjectOptions{})
}

func (m *MinioClient) RemoveObject(ctx context.Context, objectName string, opts minio.RemoveObjectOptions) error {
	fullObjectName := path.Join("attachments", objectName)
	return m.Client.RemoveObject(ctx, m.Bucket, fullObjectName, opts)
}