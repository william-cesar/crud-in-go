package email

import "github.com/william-cesar/crud-in-go/src/config/ierrors"

func NewEmailService() IEmailService {
	return &tEmailService{}
}

type IEmailService interface {
	SendActivationEmail(email, subject, template string) *ierrors.TError
}

type tEmailService struct{}
