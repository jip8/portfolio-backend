package services

import (
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient(cfg *entity.Config) (*minio.Client, error) {
	endpoint := cfg.Minio.Endpoint
	accessKeyID := cfg.Minio.AccessKeyID
	secretAccessKey := cfg.Minio.SecretAccessKey
	useSSL := cfg.Minio.UseSSL

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
