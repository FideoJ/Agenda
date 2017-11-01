package err

import "errors"
import "strings"

var (
	UserAlreadyExists         error = errors.New("User already exists")
	MeetingTitleAlreadyExists error = errors.New("Title already exists")
	WrongUsernameOrPassword   error = errors.New("Wrong username or password")
	RequireLoggedIn           error = errors.New("Not logged in yet")
)

func RequireNonEmpty(key string) error {
	return errors.New(strings.Title(key) + " must be a non-empty value")
}
