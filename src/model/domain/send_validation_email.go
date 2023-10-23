package domain

import (
	"fmt"
	"os"

	es "github.com/william-cesar/crud-in-go/src/config/email"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
	"github.com/william-cesar/crud-in-go/src/config/logger"
)

func generateActivationMailTemplate(id string) string {
	logger.NewInfoLog(logger.JOURNEY["ACTIVATE"], logger.MESSAGE["INIT"]["TEMPLATE"])

	baseUrl, version := os.Getenv("BASE_URL"), os.Getenv("ENTRYPOINT")

	activationLink := fmt.Sprintf("%s/%s/activate/%s", baseUrl, version, id)

	templateStr := fmt.Sprintf(`
		Hello! Thank your for subscribing to Users API. <br/>
		To activate your account, click on the link below: <br/>
		<a href="%s" target="_blank">Activate account</a>`, activationLink,
	)

	logger.NewInfoLog(logger.JOURNEY["ACTIVATE"], logger.MESSAGE["OK"]["TEMPLATE"])
	return templateStr
}

func (u *tUser) SendActivationEmail(email, id string) *ierrors.TError {
	logger.NewInfoLog(logger.JOURNEY["ACTIVATE"], logger.MESSAGE["INIT"]["SEND"])

	subject := "Activate your account"
	template := generateActivationMailTemplate(id)

	if err := es.NewEmailService().SendActivationEmail(email, subject, template); err != nil {
		logger.NewErrorLog(logger.JOURNEY["ACTIVATE"], logger.MESSAGE["ERROR"]["SEND_EMAIL"], err)
		return err
	}

	logger.NewInfoLog(logger.JOURNEY["ACTIVATE"], logger.MESSAGE["OK"]["SENT"])
	return nil
}
