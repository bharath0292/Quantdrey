package qerrors

import "errors"

var (
	ErrAlreadyExists = errors.New("Already exists")
	ErrNotFound      = errors.New("Not found")

	ErrInternal = errors.New("Something went wrong")
)
