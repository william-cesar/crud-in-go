package adapters

import (
	"github.com/william-cesar/crud-in-go/src/model/domain"
)

func ConvertDomainToResponse(domainUser domain.IUser) *TUserResponse {
	return &TUserResponse{
		ID:     domainUser.GetId(),
		Name:   domainUser.GetName(),
		Email:  domainUser.GetEmail(),
		Age:    domainUser.GetAge(),
		Active: domainUser.IsActive(),
	}
}
