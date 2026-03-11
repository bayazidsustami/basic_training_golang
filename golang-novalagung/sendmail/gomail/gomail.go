package main

import (
	"gopkg.in/gomail.v2"
	"log"
)

const _CONFIG_SMTP_HOST = "smtp.gmail.com"
const _CONFIG_SMTP_PORT = 587
const _CONFIG_SENDER_NAME = "Bay's Corporation <jagonadev09@gmail.com>"
const _CONFIG_AUTH_EMAIL = "jagonadev09@gmail.com"
const _CONFIG_AUTH_PASSWORD = "jagona.09"

func main() {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", _CONFIG_SENDER_NAME)
	mailer.SetHeader("To", "bayazidsustamy@gmail.com", "amy_yazid@yahoo.com", "amybayazid@gmail.com")
	mailer.SetAddressHeader("Cc", "ryanrafliappe@gmail.com", "Test test lalla")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", "Hello, <b>halo from html</b>")
	mailer.Attach("./sample.jpg")

	dialer := gomail.NewDialer(
		_CONFIG_SMTP_HOST,
		_CONFIG_SMTP_PORT,
		_CONFIG_AUTH_EMAIL,
		_CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Println("Mail sent!")
}
