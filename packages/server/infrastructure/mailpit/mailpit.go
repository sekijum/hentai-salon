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
	host := os.Getenv("MAILPIT_HOST")
	port := os.Getenv("MAILPIT_PORT")

	return &MailpitClient{
		Host: host,
		Port: port,
	}
}

func (c *MailpitClient) SendEmail(to, subject, body string) error {
	from := "no-reply@example.com" // 送信者アドレスを適切に設定
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	addr := fmt.Sprintf("%s:%s", c.Host, c.Port)
	return smtp.SendMail(addr, nil, from, []string{to}, []byte(msg))
}
