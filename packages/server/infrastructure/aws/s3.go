package aws

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const S3PresignedURLDuration = 15 * time.Minute

type S3Client struct {
	Client *s3.Client
}

func NewS3Client() (*S3Client, error) {
	region := os.Getenv("AWS_REGION")
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	return &S3Client{Client: client}, nil
}

func (s *S3Client) GeneratePresignedURL(bucketName, objectName string) (string, error) {
	presignClient := s3.NewPresignClient(s.Client)
	input := &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectName,
	}
	req, err := presignClient.PresignPutObject(context.TODO(), input, s3.WithPresignExpires(S3PresignedURLDuration))
	if err != nil {
		return "", err
	}
	return req.URL, nil
}
