package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const S3PresignedURLDuration = 15 * time.Minute

type S3Client struct {
	Client *s3.Client
}

func NewS3Client() (*S3Client, error) {
	cfg, err := LoadAWSConfig()
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	return &S3Client{Client: client}, nil
}

type S3ClientGeneratePresignedURL struct {
	BucketName, ObjectName string
}

func (s *S3Client) GeneratePresignedURL(params S3ClientGeneratePresignedURL) (string, error) {
	presignClient := s3.NewPresignClient(s.Client)
	input := &s3.PutObjectInput{
		Bucket: &params.BucketName,
		Key:    &params.ObjectName,
	}
	req, err := presignClient.PresignPutObject(context.TODO(), input, s3.WithPresignExpires(S3PresignedURLDuration))
	if err != nil {
		return "", err
	}
	return req.URL, nil
}
