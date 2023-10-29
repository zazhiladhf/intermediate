package app

import (
	"mailcampaign/config"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type Service struct {
	// repo Repository
}

func NewService() Service {
	return Service{}
}

func (s Service) SendMailService(ctx *gin.Context, to []string, cc []string, subject string, message string) (err error) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.Cfg.SMTP.SenderName)
	mailer.SetHeader("To", to...)

	for _, ccEmail := range cc {
		mailer.SetAddressHeader("Cc", ccEmail, "")
	}

	mailer.SetHeader("Subject", subject)

	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		config.Cfg.SMTP.Host,
		config.Cfg.SMTP.Port,
		config.Cfg.SMTP.AuthEmail,
		config.Cfg.SMTP.AuthPassword,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	return
}
