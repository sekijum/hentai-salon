package minio

import (
	"context"
	"os"
	"time"
	"net/http"
	"net/url"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const MinioPresignedURLDuration = 15 * time.Minute

type MinioClient struct {
	Client *minio.Client
}

func NewMinioClient() (*MinioClient, error) {
	MINIO_INTERNAL_ENDPOINT := os.Getenv("MINIO_INTERNAL_ENDPOINT")
	MINIO_EXTERNAL_ENDPOINT := os.Getenv("MINIO_EXTERNAL_ENDPOINT")
	MINIO_ROOT_USER := os.Getenv("MINIO_ROOT_USER")
	MINIO_ROOT_PASSWORD := os.Getenv("MINIO_ROOT_PASSWORD")
	useSSL := os.Getenv("MINIO_USE_SSL") == "false"

	// プロキシを設定
	proxyURL, err := url.Parse("http://" + MINIO_INTERNAL_ENDPOINT)
	if err != nil {
					return nil, err
	}
	transport := &http.Transport{
					Proxy: http.ProxyURL(proxyURL),
	}
	client, err := minio.New(MINIO_EXTERNAL_ENDPOINT, &minio.Options{
					Creds:     credentials.NewStaticV4(MINIO_ROOT_USER, MINIO_ROOT_PASSWORD, ""),
					Secure:    useSSL,
					Transport: transport,
	})
	if err != nil {
					return nil, err
	}
	return &MinioClient{Client: client}, nil
}

func (m *MinioClient) GeneratePresignedURL(bucketName, objectName string) (string, error) {
	presignedURL, err := m.Client.PresignedPutObject(context.Background(), bucketName, objectName, MinioPresignedURLDuration)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
