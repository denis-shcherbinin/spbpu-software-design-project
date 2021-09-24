package errs

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)
