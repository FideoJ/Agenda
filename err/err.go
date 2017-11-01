package err

import "errors"
import "strings"

var (
	UserAlreadyExists         error = errors.New("User already exists")
	MeetingTitleAlreadyExists error = errors.New("Title already exists")
	WrongUsernameOrPassword   error = errors.New("Wrong username or password")
)

func RequireNonEmpty(key string) error {
	return errors.New(strings.Title(key) + " must be a non-empty value")
}
