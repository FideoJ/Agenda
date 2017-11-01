package err

import "errors"

var (
	UserAlreadyExists         error = errors.New("user already exists")
	MeetingTitleAlreadyExists error = errors.New("title already exists")
	RegWithEmptyUsername error = errors.New("Register with Empty Username")
	RegWithEmptyPassword error = errors.New("Register with Empty Password")
)
