package pkg

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	Viper *viper.Viper
	Log   *logrus.Logger
}

func NewMail(viper *viper.Viper, log *logrus.Logger) *Mail {
	return &Mail{
		Viper: viper,
		Log:   log,
	}
}

// SendMail implements MailService.
func (m *Mail) SendMail(to []string, subject string, body string) error {
	username := m.Viper.GetString("mail.username")
	smtpHost := m.Viper.GetString("mail.host")
	smtpPort := m.Viper.GetInt("mail.port")
	password := m.Viper.GetString("mail.password")

	if username == "" {
		return fmt.Errorf("email sender address is empty")
	}

	if len(to) == 0 {
		return fmt.Errorf("recipient address is empty")
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", username)
	mail.SetHeader("To", to...)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	d := gomail.NewDialer(smtpHost, smtpPort, username, password)

	return d.DialAndSend(mail)
}
