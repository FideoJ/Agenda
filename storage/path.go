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

func CreateAgendaDir() {
	path := AgendaDir()
	mode := os.FileMode(0755)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
	}
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
