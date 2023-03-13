package transformation

import (
	"time"
	serviceModel "user-api/service/model"
	"user-api/transport/model/dto"
)

func ToServiceModel(user dto.User) serviceModel.User {
	return serviceModel.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		BirthDate: serviceModel.DateTime{
			Day:   user.BirthDate.Day,
			Month: time.Month(user.BirthDate.Month),
			Year:  user.BirthDate.Year,
		},
	}
}

func ToDtoUser(user serviceModel.User) dto.User {
	return dto.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		BirthDate: dto.BirthDate{
			Day:   user.BirthDate.Day,
			Month: int(user.BirthDate.Month),
			Year:  user.BirthDate.Year,
		},
	}
}

func ToDtoUsers(users []serviceModel.User) []dto.User {
	dtoUsers := make([]dto.User, 0, len(users))
	for _, user := range users {
		dtoUsers = append(dtoUsers, ToDtoUser(user))
	}
	return dtoUsers
}
