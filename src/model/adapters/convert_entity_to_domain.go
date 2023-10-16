package adapters

import (
	d "github.com/william-cesar/crud-in-go/src/model/domain"
	e "github.com/william-cesar/crud-in-go/src/model/entity"
)

func ConvertEntityToDomain(ue e.TuserEntity) d.IUser {
	domain := d.NewUser(
		ue.Email,
		ue.Password,
		ue.Name,
		ue.Age,
	)

	domain.SetID(ue.ID.Hex())
	return domain
}
