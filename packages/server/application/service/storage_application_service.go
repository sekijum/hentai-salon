package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	aws "server/infrastructure/aws"
	minio "server/infrastructure/minio"
	"server/presentation/request"
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

type StorageApplicationServiceGeneratePresignedURLs struct {
	Ctx  context.Context
	Body request.GeneratePresignedURLsRequest
}

func (s *StorageApplicationService) GeneratePresignedURLs(params StorageApplicationServiceGeneratePresignedURLs) ([]string, error) {
	bucketName := os.Getenv("AWS_BUCKET_NAME")
	urls := make([]string, len(params.Body.ObjectNameList))

	for i, objectName_i := range params.Body.ObjectNameList {
		now := time.Now()
		objectKey := fmt.Sprintf("%d/%02d/%02d/%s", now.Year(), now.Month(), now.Day(), objectName_i)

		var url string
		var err error

		if os.Getenv("APP_ENV") == "production" {
			if s.s3Client == nil {
				return nil, errors.New("S3 クライアントが初期化されていません")
			}
			url, err = s.s3Client.GeneratePresignedURL(aws.S3ClientGeneratePresignedURL{
				BucketName: bucketName,
				ObjectName: objectKey,
			})
		} else {
			if s.minioClient == nil {
				return nil, errors.New("minio クライアントが初期化されていません")
			}
			url, err = s.minioClient.GeneratePresignedURL(minio.MinioClientGeneratePresignedURL{
				Ctx:        params.Ctx,
				BucketName: bucketName,
				ObjectName: objectKey,
			})
		}

		if err != nil {
			return nil, err
		}
		urls[i] = url
	}
	return urls, nil
}
