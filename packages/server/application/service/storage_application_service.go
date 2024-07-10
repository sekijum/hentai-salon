package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	aws "server/infrastructure/aws"
	minio "server/infrastructure/minio"
)

type StorageApplicationService struct {
	minioClient *minio.MinioClient
	s3Client    *aws.S3Client
}

func NewStorageApplicationService(minioClient *minio.MinioClient, s3Client *aws.S3Client) *StorageApplicationService {
	return &StorageApplicationService{
		minioClient: minioClient,
		s3Client:    s3Client,
	}
}

func (s *StorageApplicationService) GeneratePresignedURLs(objectNames []string) ([]string, error) {
	env := os.Getenv("APP_ENV")
	bucketName := os.Getenv("AWS_BUCKET_NAME")
	urls := make([]string, len(objectNames))

	for i, objectName := range objectNames {
		now := time.Now()
		objectKey := fmt.Sprintf("%d/%02d/%02d/%s", now.Year(), now.Month(), now.Day(), objectName)

		var url string
		var err error

		if env == "development" {
			if s.minioClient == nil {
				return nil, errors.New("Minio クライアントが初期化されていません")
			}
			url, err = s.minioClient.GeneratePresignedURL(bucketName, objectKey)
		} else if env == "production" {
			if s.s3Client == nil {
				return nil, errors.New("S3 クライアントが初期化されていません")
			}
			url, err = s.s3Client.GeneratePresignedURL(bucketName, objectKey)
		} else {
			return nil, errors.New("不明な環境設定です")
		}

		if err != nil {
			return nil, err
		}
		urls[i] = url
	}
	return urls, nil
}
