package err

import "errors"

var (
	UserAlreadyExists error = errors.New("user already exists")
)
