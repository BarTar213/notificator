package email

import (
	"fmt"
	"time"

	"github.com/BarTar213/notificator/config"
	"github.com/xhit/go-simple-mail/v2"
)

type Client interface {
	Send(email *mail.Email) error
}

type Email struct {
	From   string
	Server *mail.SMTPServer
	Client *mail.SMTPClient
}

func New(conf *config.Mail) (Client, error) {
	server := mail.NewSMTPClient()

	server.Host = conf.Host
	server.Port = conf.Port
	server.Username = conf.Username
	server.Password = conf.Password
	server.Encryption = mail.EncryptionTLS
	server.KeepAlive = true

	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		return nil, err
	}

	return &Email{
		From:   conf.From,
		Server: server,
		Client: smtpClient,
	}, nil
}

func (c *Email) Send(email *mail.Email) error {
	err := c.Client.Noop()
	if err != nil {
		c.Client, err = c.Server.Connect()
		if err != nil {
			return err
		}
	}

	return email.SetFrom(fmt.Sprintf("%s <%s>", c.From, c.Server.Username)).
		Send(c.Client)
}
