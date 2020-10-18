package senders

import (
	"github.com/BarTar213/notificator/email"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Email struct {
	TemplateName string
	Recipients   []string
	Data         map[string]string
}

func (e *Email) Send(c email.Client, subject, message string, html bool) error {
	msg := mail.NewMSG()
	msg.SetSubject(subject).
		AddTo(e.Recipients...)

	if html {
		msg.SetBody(mail.TextHTML, message)
		return c.Send(msg)
	}

	msg.SetBody(mail.TextPlain, message)
	return c.Send(msg)
}

func (e *Email) GetTemplateName() string {
	return e.TemplateName
}

func (e *Email) GetData() map[string]string {
	return e.Data
}
