package errs

import "errors"

var (
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidUsernameLength = errors.New("invalid username length")
	ErrInvalidPasswordLength = errors.New("invalid password length")
)
