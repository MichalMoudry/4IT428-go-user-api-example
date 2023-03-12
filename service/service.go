package service

import (
	"context"
	"user-api/service/errors"
	"user-api/service/model"
)

type Service struct{}

func CreateService() Service {
	return Service{}
}

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

func (Service) ListUsers(_ context.Context) []model.User {
	usersList := make([]model.User, 0, len(users))
	for _, user := range users {
		usersList = append(usersList, user)
	}

	return usersList
}

func (Service) GetUser(_ context.Context, email string) (model.User, error) {
	user, exists := users[email]
	if !exists {
		return model.User{}, errors.ErrUserDoesntExist
	}
	return user, nil
}

func (Service) UpdateUser(_ context.Context, email string, user model.User) (model.User, error) {
	origUser, exists := users[email]
	if !exists {
		return model.User{}, errors.ErrUserDoesntExist
	}

	if origUser.Email == email {
		users[email] = user
	} else {
		users[user.Email] = user
		delete(users, email)
	}

	return user, nil
}

func (Service) DeleteUser(_ context.Context, email string) error {
	_, exists := users[email]
	if !exists {
		return errors.ErrUserDoesntExist
	}
	delete(users, email)
	return nil
}
