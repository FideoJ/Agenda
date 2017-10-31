package err

import "errors"

var (
	UserAlreadyExists error = errors.New("user already exists")
	UserNotExists error = errors.New("user doesn't exists")
	AlreadyLogin error = errors.New("already login")
)
