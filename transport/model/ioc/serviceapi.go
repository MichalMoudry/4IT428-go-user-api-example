package ioc

import (
	"context"
	serviceModel "user-api/service/model"
)

type IService interface {
	CreateUser(_ context.Context, user serviceModel.User) error
	ListUsers(_ context.Context) []serviceModel.User
	GetUser(_ context.Context, email string) (serviceModel.User, error)
	UpdateUser(_ context.Context, email string, user serviceModel.User) (serviceModel.User, error)
	DeleteUser(_ context.Context, email string) error
}
