package library

import (
	"log"
	"net/smtp"

	"../config"
)

func SendMail(to string, subject string, body string){
	from := config.Mail_username
	pass := config.Mail_Password

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";
	subject = "Subject: " + subject + "\n"
	msg := []byte(subject + mime + body)

	/*msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body */

	err := smtp.SendMail(config.Mail_smtp_addr,
		smtp.PlainAuth("", from, pass, config.Mail_smtp_host),
		from, []string{to}, msg)

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}