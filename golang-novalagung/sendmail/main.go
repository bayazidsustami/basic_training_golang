package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Bay's Corporation <jagonadev09@gmail.com>"
const CONFIG_AUTH_EMAIL = "jagonadev09@gmail.com"
const CONFIG_AUTH_PASSWORD = "jagona.09"

func main() {
	to := []string{"bayazidsustamy@gmail.com", "amy_yazid@yahoo.com", "amybayazid@gmail.com"}
	cc := []string{"ryanrafliappe@gmail.com"}
	subject := "Test send email from golang"
	message := "hello"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("mail sent!")
}

func sendMail(to []string, cc []string, subject string, message string) error {
	body := "From : " + CONFIG_SENDER_NAME + "\n" +
		"To : " + strings.Join(to, ",") + "\n" +
		"Cc : " + strings.Join(cc, ",") + "\n" +
		"Subject : " + subject + "\n\n" +
		message
	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
