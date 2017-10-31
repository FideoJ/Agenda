package err

import "errors"

var (
	UserAlreadyExists         error = errors.New("user already exists")
	MeetingTitleAlreadyExists error = errors.New("title already exists")
)
