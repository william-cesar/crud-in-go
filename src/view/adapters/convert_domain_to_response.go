package adapters

import (
	"github.com/william-cesar/crud-in-go/src/model/domain"
)

func ConvertDomainToResponse(domainUser domain.IUser) *TUserResponse {
	return &TUserResponse{
		ID:    "",
		Name:  domainUser.GetName(),
		Email: domainUser.GetEmail(),
		Age:   domainUser.GetAge(),
	}
}
