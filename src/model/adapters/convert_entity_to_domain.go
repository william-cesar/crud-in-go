package adapters

import (
	"github.com/william-cesar/crud-in-go/src/model/domain"
	"github.com/william-cesar/crud-in-go/src/model/entity"
)

func ConvertEntityToDomain(ue entity.TuserEntity) domain.IUser {
	user := domain.NewUser(
		ue.Email,
		ue.Password,
		ue.Name,
		ue.Age,
		ue.Active,
	)

	user.SetID(ue.ID.Hex())
	return user
}
