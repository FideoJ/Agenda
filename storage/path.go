package storage

import (
	"os"
)

func AgendaDir() string {
	home, present := os.LookupEnv("HOME")
	if !present {
		home = "."
	}
	return home + "/.agenda/"
}

func SessionFile() string {
	return AgendaDir() + "session.json"
}

func UserFile() string {
	return AgendaDir() + "users.json"
}

func MeetingFile() string {
	return AgendaDir() + "meetings.json"
}

func CreateAgendaDir() {
	os.Mkdir(AgendaDir(), 0755)
}

func RemoveSessionFile() {
	os.Remove(SessionFile())
}
