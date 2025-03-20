package mailer

import (
	"go-training/validation-api/configs"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type SendEmailHandler struct {
	configs.SmtpConfig
}

func NewSendEmailHandler(SmtpConfig configs.SmtpConfig) *SendEmailHandler {
	return &SendEmailHandler{
		SmtpConfig,
	}
}

func (handler *SendEmailHandler) Send(emailAddress string, subject string, html string) (bool, error) {
	e := email.NewEmail()
	e.From = `Test <` + handler.SmtpEmail + `>`
	e.To = []string{emailAddress}
	e.Subject = subject
	e.HTML = []byte(html)
	err := e.Send(handler.SmtpAddress+":587", smtp.PlainAuth("", handler.SmtpEmail, handler.SmtpPassword, handler.SmtpAddress))
	if err != nil {
		return false, err
	}
	return true, nil
}
