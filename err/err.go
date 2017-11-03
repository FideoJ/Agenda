package err

import (
	"errors"
	"fmt"
	"strings"

	"../entity"
)

var (
	UserAlreadyExists         error = errors.New("User already exists")
	WrongUsernameOrPassword   error = errors.New("Wrong username or password")
	RequireLoggedIn           error = errors.New("Not logged in yet")
	UserNotExist              error = errors.New("User not found")
	AttendantsDuplicated      error = errors.New("Attendants duplicated")
	TitleAlreadyExists        error = errors.New("Title already exists")
	StartTimeNotBeforeEndTime error = errors.New("Start time not before end time")
	MeetingNotFound           error = errors.New("Meeting not found")
	SponsorRemoveSelf         error = errors.New("Sponsor remove himself")
)

func RequireNonEmpty(key string) error {
	return errors.New(strings.Title(key) + " must be a non-empty value")
}

func TimeConflicted(attendant string, relatedMeeting *entity.Meeting) error {
	const timeLayout = "2006-01-02 15:04"
	sTimeStr := relatedMeeting.StartTime.Format(timeLayout)
	eTimeStr := relatedMeeting.EndTime.Format(timeLayout)
	text := fmt.Sprintf("%s encounters a time conflict because of "+
		"another meeting: title:[%s], startTime:[%s], endTime[%s]",
		strings.Title(attendant), relatedMeeting.Title, sTimeStr, eTimeStr)

	return errors.New(text)
}
