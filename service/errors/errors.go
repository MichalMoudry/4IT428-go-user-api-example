package errors

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserDoesntExist   = errors.New("user does not exist")
)
