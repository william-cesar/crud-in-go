package email

import (
	"os"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	gomail "gopkg.in/mail.v2"
)

const (
	host = "smtp.gmail.com"
	port = 587
)

func (es *tEmailService) SendActivationEmail(email, subject, template string) *ierrors.TError {
	from, password := os.Getenv("EMAIL_SENDER"), os.Getenv("EMAIL_PASSWORD")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", template)

	d := gomail.NewDialer(host, port, from, password)

	if err := d.DialAndSend(m); err != nil {
		return ierrors.NewInternalError(err.Error())
	}

	return nil
}
