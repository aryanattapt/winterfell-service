package pkg

import (
	"errors"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type MailPayload struct {
	To      []string
	Cc      []string
	Subject string
	Message string
}

func (payload MailPayload) SendMail() (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_EMAIL"))
	if len(payload.To) == 0 {
		err = errors.New("destination is empty")
		return
	} else {
		m.SetHeader("To", payload.To...)
	}
	if len(payload.Cc) != 0 {
		m.SetHeader("Cc", payload.Cc...)
	}
	m.SetHeader("Subject", payload.Subject)
	m.SetBody("text/html", payload.Message)

	smtp_port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), smtp_port, os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"))
	return d.DialAndSend(m)
}
