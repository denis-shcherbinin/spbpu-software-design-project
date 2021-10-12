package errs

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrListNotFound      = errors.New("todo-list with such id not found")
	ErrItemNotFound      = errors.New("todo-item with such id not found")
)
