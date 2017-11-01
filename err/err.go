package err

import "errors"

var (
	RegWithEmptyUsername error = errors.New("Register with Empty Username")
	RegWithEmptyPassword error = errors.New("Register with Empty Password")
	UserAlreadyExists    error = errors.New("user already exists")
)
