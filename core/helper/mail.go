package helper

import (
	"go-commerce/config"
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendMail(toEmail string, subject string, body string) error {

	host := config.Config("MAIL_HOST")
	port, _ := strconv.Atoi(config.Config("MAIL_PORT"))
	username := config.Config("MAIL_USER")
	password := config.Config("MAIL_PASSWORD")

	from := config.Config("MAIL_FROM")

	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", toEmail)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	dialer := gomail.NewDialer(host, port, username, password)

	err := dialer.DialAndSend(message)

	if err != nil {
		return err
	}

	return nil
}
