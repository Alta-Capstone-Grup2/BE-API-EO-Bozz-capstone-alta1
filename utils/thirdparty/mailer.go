package thirdparty

import (
	"bytes"
	"capstone-alta1/utils/helper"
	"errors"
	"fmt"
	"os"
	"path"
	"text/template"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "EO-Bozz <eobozz01@gmail.com>"

func SendMail(recipientEmail, subject string, data interface{}, fileTemplate string) error {

	emailBody, errParseTempalte := ParseTemplate(fileTemplate, data)
	if errParseTempalte != nil {
		helper.LogDebug("mailer - SendMail - Failed ParseTemplate. Error ", errParseTempalte)
		return errors.New("Failed sent email. Please try again.")
	}

	CONFIG_AUTH_EMAIL := os.Getenv("MAILER_SENDER_EMAIL")
	CONFIG_AUTH_PASSWORD := os.Getenv("MAILER_SENDER_PASSWORD")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", recipientEmail)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", emailBody)
	// mailer.Attach("https://project3bucker.s3.ap-southeast-1.amazonaws.com/partner/EEz06AIRAiyJe4ghKfU5-default_image.jpg")

	dialer := &gomail.Dialer{
		Host:     CONFIG_SMTP_HOST,
		Port:     CONFIG_SMTP_PORT,
		Username: CONFIG_AUTH_EMAIL,
		Password: CONFIG_AUTH_PASSWORD,
	}

	helper.LogDebug("Gomail email recipient ", recipientEmail)
	helper.LogDebug("Gomail config ", CONFIG_AUTH_EMAIL, " pass", CONFIG_AUTH_PASSWORD)
	helper.LogDebug("Gomail dialer data", *dialer)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		helper.LogDebug("Failed sent email. Error : ", err)
		return errors.New("Failed sent email. Please try again.")
	}
	helper.LogDebug("Success sent email. ")

	return nil
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		helper.LogDebug("mailer - ParseTempalte - Failed ParseTemplate. Error ", err)
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		helper.LogDebug("mailer - ParseTempalte - Failed Execute. Error ", err)
		return "", err
	}
	return buf.String(), nil
}

func SendEmailWaitingPayment2(recipientEmail, subject string, data interface{}) {

	template := path.Join("utils", "thirdparty", "payment_waiting_for_payment.html")
	err := SendMail(recipientEmail, subject, data, template)
	if err != nil {
		helper.LogDebug("SendEmailWaitingPayment2 - Failed send email")
		fmt.Println("send email '" + subject + "' success")
	} else {
		fmt.Println(err)
	}
}

func SendMailWaitingPayment(emailClient string) {
	CONFIG_AUTH_EMAIL := os.Getenv("MAILER_SENDER_EMAIL")
	CONFIG_AUTH_PASSWORD := os.Getenv("MAILER_SENDER_PASSWORD")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailClient)
	mailer.SetHeader("Subject", "Waiting for Payment from EO-Bozz")
	mailer.SetBody("text/html", "Hello World!, <p>Waiting for Payment, please payout your transaction.</p>")
	// mailer.Attach("https://project3bucker.s3.ap-southeast-1.amazonaws.com/partner/EEz06AIRAiyJe4ghKfU5-default_image.jpg")

	dialer := &gomail.Dialer{
		Host:     CONFIG_SMTP_HOST,
		Port:     CONFIG_SMTP_PORT,
		Username: CONFIG_AUTH_EMAIL,
		Password: CONFIG_AUTH_PASSWORD,
	}

	helper.LogDebug("Gomail email recipient ", emailClient)
	helper.LogDebug("Gomail config ", CONFIG_AUTH_EMAIL, " pass", CONFIG_AUTH_PASSWORD)
	helper.LogDebug("Gomail dialer data", *dialer)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		helper.LogDebug("Failed sent email. Error : ", err)
	} else {
		helper.LogDebug("Success sent email. ")
	}
}

func SendMailConfirmedOrder(emailClient, Calendar string) {
	CONFIG_AUTH_EMAIL := os.Getenv("MAILER_SENDER_EMAIL")
	CONFIG_AUTH_PASSWORD := os.Getenv("MAILER_SENDER_PASSWORD")
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailClient)
	mailer.SetHeader("Subject", "Confirmed Email from EO-Bozz")
	mailer.SetBody("text/html", "Hello World!, <p>Thanks for Order, see you later. :) </p>"+"<div>"+Calendar+"</div>")
	mailer.Attach(Calendar)

	dialer := &gomail.Dialer{
		Host:     CONFIG_SMTP_HOST,
		Port:     CONFIG_SMTP_PORT,
		Username: CONFIG_AUTH_EMAIL,
		Password: CONFIG_AUTH_PASSWORD,
	}

	helper.LogDebug("Gomail email recipient ", emailClient)
	// helper.LogDebug("Gomail config ", CONFIG_AUTH_EMAIL, " pass", CONFIG_AUTH_PASSWORD)
	// helper.LogDebug("Gomail dialer data", *dialer)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		helper.LogDebug("Failed sent email. Error : ", err)
	} else {
		helper.LogDebug("Success sent email. ")
	}
}

func SendMailWaitingConfirmation(emailClient string) {
	CONFIG_AUTH_EMAIL := os.Getenv("MAILER_SENDER_EMAIL")
	CONFIG_AUTH_PASSWORD := os.Getenv("MAILER_SENDER_PASSWORD")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailClient)
	mailer.SetHeader("Subject", "Payout Success, please wait your EO Confirmation")
	mailer.SetBody("text/html", "Hello World!, <p>Thanks for your paid, your EO is checking payment. Please wait for confirmation from EO.</p>")
	// mailer.Attach("https://project3bucker.s3.ap-southeast-1.amazonaws.com/partner/EEz06AIRAiyJe4ghKfU5-default_image.jpg")

	dialer := &gomail.Dialer{
		Host:     CONFIG_SMTP_HOST,
		Port:     CONFIG_SMTP_PORT,
		Username: CONFIG_AUTH_EMAIL,
		Password: CONFIG_AUTH_PASSWORD,
	}

	helper.LogDebug("Gomail email recipient ", emailClient)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		helper.LogDebug("Failed sent email. Error : ", err)
	} else {
		helper.LogDebug("Success sent email. ")
	}
}

func SendMailCompleteOrder(emailClient string) {
	CONFIG_AUTH_EMAIL := os.Getenv("MAILER_SENDER_EMAIL")
	CONFIG_AUTH_PASSWORD := os.Getenv("MAILER_SENDER_PASSWORD")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailClient)
	mailer.SetHeader("Subject", "Order Done")
	mailer.SetBody("text/html", "Hello World!, <p>Thank you for your cooperation, see you next time.</p>")
	// mailer.Attach("https://project3bucker.s3.ap-southeast-1.amazonaws.com/partner/EEz06AIRAiyJe4ghKfU5-default_image.jpg")

	dialer := &gomail.Dialer{
		Host:     CONFIG_SMTP_HOST,
		Port:     CONFIG_SMTP_PORT,
		Username: CONFIG_AUTH_EMAIL,
		Password: CONFIG_AUTH_PASSWORD,
	}

	helper.LogDebug("Gomail email recipient ", emailClient)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		helper.LogDebug("Failed sent email. Error : ", err)
	} else {
		helper.LogDebug("Success sent email. ")
	}
}

func SendMailPayoutSuccess(emailClient string) {
	CONFIG_AUTH_EMAIL := os.Getenv("MAILER_SENDER_EMAIL")
	CONFIG_AUTH_PASSWORD := os.Getenv("MAILER_SENDER_PASSWORD")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailClient)
	mailer.SetHeader("Subject", "Payout Done")
	mailer.SetBody("text/html", "Hello World!, <p>Thank you for using EO-Bozz as your business advice.</p>")
	// mailer.Attach("https://project3bucker.s3.ap-southeast-1.amazonaws.com/partner/EEz06AIRAiyJe4ghKfU5-default_image.jpg")

	dialer := &gomail.Dialer{
		Host:     CONFIG_SMTP_HOST,
		Port:     CONFIG_SMTP_PORT,
		Username: CONFIG_AUTH_EMAIL,
		Password: CONFIG_AUTH_PASSWORD,
	}

	helper.LogDebug("Gomail email recipient ", emailClient)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		helper.LogDebug("Failed sent email. Error : ", err)
	} else {
		helper.LogDebug("Success sent email. ")
	}
}
