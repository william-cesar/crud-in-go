package adapters

import (
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"github.com/william-cesar/crud-in-go/src/model/entity"
)

func ConvertDomainToEntity(
	ud domain.IUser,
) *entity.TuserEntity {
	return &entity.TuserEntity{
		Email:    ud.GetEmail(),
		Password: ud.GetPassword(),
		Name:     ud.GetName(),
		Age:      ud.GetAge(),
		Active:   ud.IsActive(),
	}
}
