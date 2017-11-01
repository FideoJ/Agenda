package entity

import (
	"encoding/json"
	"io"
	"time"

	"../logger"
)

type Meeting struct {
	Title        string
	StartTime    time.Time
	EndTime      time.Time
	Sponsor      *User
	Participants []string
}

type Meetings map[string]*Meeting

func (meetings Meetings) Query(title string) *Meeting {
	return meetings[title]
}

func (meetings Meetings) Add(meeting *Meeting) {
	meetings[meeting.Title] = meeting
}

func (meetings Meetings) Serialize(w io.Writer) {
	encoder := json.NewEncoder(w)
	var err error

	for _, meeting := range meetings {
		err = encoder.Encode(meeting)
		logger.FatalIf(err)
	}
}

func DeserializeMeeting(r io.Reader) Meetings {
	decoder := json.NewDecoder(r)
	meetings := make(Meetings)
	var meeting Meeting
	var err error

	for {
		err = decoder.Decode(&meeting)
		if err == io.EOF {
			return meetings
		}
		if err != nil {
			logger.FatalIf(err)
		}
		meetings.Add(&meeting)
	}
}
