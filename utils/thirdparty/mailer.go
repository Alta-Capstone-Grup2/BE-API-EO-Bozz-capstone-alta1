package thirdparty

import (
	"log"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "EO-Bozz <eobozz01@gmail.com>"
const CONFIG_AUTH_EMAIL = "eobozz01@gmail.com"
const CONFIG_AUTH_PASSWORD = "poiuy09876"

func SendMail(emailClient string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailClient)
	mailer.SetAddressHeader("Cc", "eobozz01@gmail.com", "EO-Bozz")
	mailer.SetHeader("Subject", "Confirmed Email from EO-Bozz")
	mailer.SetBody("text/html", "Hello, <p>Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. </p>")
	mailer.Attach("https://project3bucker.s3.ap-southeast-1.amazonaws.com/partner/EEz06AIRAiyJe4ghKfU5-default_image.jpg")

	dialer := &gomail.Dialer{
		Host:     CONFIG_SMTP_HOST,
		Port:     CONFIG_SMTP_PORT,
		Username: CONFIG_AUTH_EMAIL,
		Password: CONFIG_AUTH_PASSWORD,
	}

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
