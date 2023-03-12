package service

import (
	"context"
	"user-api/service/errors"
	"user-api/service/model"
)

var (
	users = make(map[string]model.User)
)

func (Service) CreateUser(_ context.Context, user model.User) error {
	if _, exists := users[user.Email]; exists {
		return errors.ErrUserAlreadyExists
	}
	users[user.Email] = user
	return nil
}
