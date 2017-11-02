package storage

import (
	"os"

	"github.com/MarshallW906/Agenda/entity"
	"github.com/MarshallW906/Agenda/logger"
)

func LoadUsers() entity.Users {
	file, err := os.OpenFile(UserFile(), os.O_RDONLY|os.O_CREATE, 0644)
	logger.FatalIf(err)
	users := entity.DeserializeUser(file)

	return users
}

func StoreUsers(users entity.Users) {
	file, err := os.OpenFile(UserFile(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	logger.FatalIf(err)
	users.Serialize(file)
}

func LoadMeetings() entity.Meetings {
	file, err := os.OpenFile(MeetingFile(), os.O_RDONLY|os.O_CREATE, 0644)
	logger.FatalIf(err)
	meetings := entity.DeserializeMeeting(file)

	return meetings
}

func StoreMeetings(meetings entity.Meetings) {
	file, err := os.OpenFile(MeetingFile(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	logger.FatalIf(err)
	meetings.Serialize(file)
}

func LoadCurUser() (string, bool) {
	file, err := os.OpenFile(SessionFile(), os.O_RDONLY|os.O_CREATE, 0644)
	logger.FatalIf(err)
	session := entity.DeserializeSession(file)

	return session.CurrentUser, session.CurrentUser != ""
}

func StoreSession(session *entity.Session) {
	file, err := os.OpenFile(SessionFile(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	logger.FatalIf(err)
	session.Serialize(file)
}
