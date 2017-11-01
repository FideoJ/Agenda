package storage

import (
	"os"

	"github.com/FideoJ/Agenda/entity"
)

func LoadUsers() entity.Users {
	file, _ := os.OpenFile("/tmp/users.json", os.O_RDONLY|os.O_CREATE, 0644)
	users := entity.DeserializeUser(file)
	return users
}

func StoreUsers(users entity.Users) {
	file, _ := os.OpenFile("/tmp/users.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	users.Serialize(file)
}

func LoadMeetings() entity.Meetings {
	file, _ := os.OpenFile("/tmp/meetings.json", os.O_RDONLY|os.O_CREATE, 0644)
	meetings := entity.DeserializeMeeting(file)
	return meetings
}

func StoreMeetings(users entity.Meetings) {
	file, _ := os.OpenFile("/tmp/meetings.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	users.Serialize(file)
}
