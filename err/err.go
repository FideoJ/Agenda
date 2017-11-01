package err

import "errors"

var (
	UserAlreadyExists         error = errors.New("user already exists")
	MeetingTitleAlreadyExists error = errors.New("title already exists")
	WrongUsernameOrPassword error = errors.New("wrong username or password")
)

func RequireNonEmpty(key string) error {
	return errors.New(key + " must be a non-empty value")
}
