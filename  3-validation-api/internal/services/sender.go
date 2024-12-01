package services

import (
	"fmt"
	"github.com/SemenVologdin/go-advanced/config"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Sender struct {
	cfg config.Mail
}

func newSenderService(cfg config.Mail) *Sender {
	return &Sender{
		cfg: cfg,
	}
}

func (service *Sender) Send(emailAddr string, hashEmail string) error {
	url := fmt.Sprintf("%s/%s/%s", "localhost:8081", "verify", hashEmail)

	e := email.NewEmail()
	e.From = service.cfg.Email
	e.To = []string{emailAddr}
	e.Subject = "Подтверждение письма"
	e.Text = []byte(fmt.Sprintf("Подтвердите емейл перейдя по ссылке: %s", url))

	return e.Send(
		service.cfg.Address,
		smtp.PlainAuth("", service.cfg.Email, service.cfg.Password, service.cfg.Host),
	)
}
