package services

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
)

func SendEmailUUID(email string, UUID string) {
	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("EMAIL_LOGIN"))

	m.SetHeader("To", email)

	m.SetHeader("Subject", "Подтверждение электронной почты University Web")

	m.SetBody("text/plain", "https://"+os.Getenv("EMAIL_HTTP")+"/api/activate/"+UUID)

	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_LOGIN"), os.Getenv("EMAIL_PASSWORD"))

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}

	return
}
