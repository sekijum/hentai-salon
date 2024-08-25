package mailpit

import (
	"fmt"
	"net/smtp"
	"os"
)

type MailpitClient struct {
	Host string
	Port string
}

func NewMailpitClient() *MailpitClient {
	if os.Getenv("APP_ENV") == "production" {
		return nil
	}

	host := os.Getenv("MAILPIT_HOST")
	port := os.Getenv("MAILPIT_PORT")

	return &MailpitClient{
		Host: host,
		Port: port,
	}
}

type MailpitClientSendEmail struct {
	To, Subject, Body string
}

func (c *MailpitClient) SendEmail(params MailpitClientSendEmail) error {
	mailFromAddress := os.Getenv("MAIL_FROM_ADDRESS")
	from := mailFromAddress
	msg := "From: " + from + "\n" +
		"To: " + params.To + "\n" +
		"Subject: " + params.Subject + "\n\n" +
		params.Body

	addr := fmt.Sprintf("%s:%s", c.Host, c.Port)
	return smtp.SendMail(addr, nil, from, []string{params.To}, []byte(msg))
}
