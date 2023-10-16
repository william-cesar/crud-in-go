package adapters

import (
	d "github.com/william-cesar/crud-in-go/src/model/domain"
	e "github.com/william-cesar/crud-in-go/src/model/entity"
)

func ConvertDomainToEntity(
	ud d.IUser,
) *e.TuserEntity {
	return &e.TuserEntity{
		Email:    ud.GetEmail(),
		Password: ud.GetPassword(),
		Name:     ud.GetName(),
		Age:      ud.GetAge(),
		Active:   ud.IsActive(),
	}
}
