package service

import (
	"fmt"
	"strconv"

	"github.com/RomanshkVolkov/test-api/internal/adapters/repository"
	"gopkg.in/gomail.v2"
)

type MailOptions struct {
	To      []string
	Subject string
	Body    string
}

// outlook
// MAILER_SMTP_HOST=smtp.office365.com
// MAILER_SMTP_PORT=587
func SendMail(options *MailOptions) (bool, error) {
	Email := repository.GetEnv("MAILER_EMAIL")
	Password := repository.GetEnv("MAILER_PASSWORD")
	SMTPHost := repository.GetEnv("MAILER_SMTP_HOST")
	SMTPPort := repository.GetEnv("MAILER_SMTP_PORT")
	port, err := strconv.Atoi(SMTPPort)
	if err != nil {
		port = 587
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "No-Reply <"+Email+">")
	m.SetHeader("To", options.To...)
	m.SetHeader("Subject", options.Subject)
	m.SetBody("text/html", options.Body)

	fmt.Printf("Sending email...")
	fmt.Println("Host: ", SMTPHost)
	fmt.Println("Port: ", port)

	d := gomail.NewDialer(SMTPHost, port, Email, Password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Print(err.Error())
		fmt.Printf("Error sending email: %v", err)
		return false, err
	}

	return true, nil
}
