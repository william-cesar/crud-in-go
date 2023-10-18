package domain

import (
	"fmt"
	"os"

	es "github.com/william-cesar/crud-in-go/src/config/email"
	"github.com/william-cesar/crud-in-go/src/config/ierrors"
)

func generateActivationMailTemplate(id string) string {
	baseUrl, version := os.Getenv("BASE_URL"), os.Getenv("ENTRYPOINT")

	activationLink := fmt.Sprintf("%s/%s/activate/%s", baseUrl, version, id)

	templateStr := fmt.Sprintf(`
		Hello! Thank your for subscribing to Users API. <br/>
		To activate your account, click on the link below: <br/>
		<a href="%s" target="_blank">Activate account</a>`, activationLink,
	)

	return templateStr
}

func (u *tUser) SendActivationEmail(email, id string) *ierrors.TError {
	subject := "Activate your account"
	template := generateActivationMailTemplate(id)

	if err := es.NewEmailService().SendActivationEmail(email, subject, template); err != nil {
		return err
	}

	return nil
}
