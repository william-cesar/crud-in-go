package email

import (
	"os"

	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
	gomail "gopkg.in/mail.v2"
)

const (
	host = "smtp.gmail.com"
	port = 587
)

func (es *tEmailService) SendActivationEmail(email, subject, template string) *ierrors.TError {
	logger.NewInfoLog(logger.JOURNEY["ACTIVATE_SERVICE"], logger.MESSAGE["INIT"]["EMAIL"])

	from, password := os.Getenv("EMAIL_SENDER"), os.Getenv("EMAIL_PASSWORD")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", template)

	d := gomail.NewDialer(host, port, from, password)

	if err := d.DialAndSend(m); err != nil {
		logger.NewErrorLog(logger.JOURNEY["ACTIVATE_SERVICE"], logger.MESSAGE["ERROR"]["SEND_EMAIL"], err)
		return ierrors.NewInternalError()
	}

	logger.NewInfoLog(logger.JOURNEY["ACTIVATE_SERVICE"], logger.MESSAGE["OK"]["SENT"])
	return nil
}
