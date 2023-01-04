package thirdparty

import (
	"capstone-alta1/utils/helper"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "EO-Bozz <eobozz01@gmail.com>"
const CONFIG_AUTH_EMAIL = "eobozz01@gmail.com"
const CONFIG_AUTH_PASSWORD = "vodofmmlitkoyieb"

func SendMail(emailClient string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailClient)
	mailer.SetAddressHeader("Cc", "eobozz01@gmail.com", "EO-Bozz")
	mailer.SetHeader("Subject", "Confirmed Email from EO-Bozz")
	mailer.SetBody("text/html", "Hello World!, <p>Thanks for Order, see you later. :) </p>")
	// mailer.Attach("https://project3bucker.s3.ap-southeast-1.amazonaws.com/partner/EEz06AIRAiyJe4ghKfU5-default_image.jpg")

	dialer := &gomail.Dialer{
		Host:     CONFIG_SMTP_HOST,
		Port:     CONFIG_SMTP_PORT,
		Username: CONFIG_AUTH_EMAIL,
		Password: CONFIG_AUTH_PASSWORD,
	}

	helper.LogDebug("Gomail email recipient ", emailClient)
	helper.LogDebug("Gomail dialer data", *dialer)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		helper.LogDebug("Failed sent email. Error : ", err)
	} else {
		helper.LogDebug("Success sent email. ")
	}
}
