package aws

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type SESClient struct {
	Client *ses.Client
}

func NewSESClient() (*SESClient, error) {
	if os.Getenv("APP_ENV") != "production" {
		return nil, nil
	}

	cfg, err := LoadAWSConfig()
	if err != nil {
		return nil, err
	}
	client := ses.NewFromConfig(cfg)
	return &SESClient{Client: client}, nil
}

type SESClientSendEmail struct {
	To, Subject, Body string
}

func (s *SESClient) SendEmail(params SESClientSendEmail) error {
	fmt.Print("-----------------------------------------------------")
	fmt.Print(params)
	fmt.Print("-----------------------------------------------------")
	input := &ses.SendEmailInput{
		Source: aws.String("変態サロン <" + os.Getenv("MAIL_FROM_ADDRESS") + ">"),
		Destination: &types.Destination{
			ToAddresses: []string{params.To},
		},
		Message: &types.Message{
			Subject: &types.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(params.Subject),
			},
			Body: &types.Body{
				Text: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(params.Body),
				},
			},
		},
	}

	_, err := s.Client.SendEmail(context.TODO(), input)
	return err
}
