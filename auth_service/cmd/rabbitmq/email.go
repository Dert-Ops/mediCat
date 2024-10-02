package rabbitmq

import (
	"net/smtp"
)

func SendEmail(to string, subject string, body string) error {
	from := "berkay_inam@hotmail.com"
	password := "1597538246bB."

	smtpHost := "smtp.office365.com"
	smtpPort := "587"

	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	return err
}
