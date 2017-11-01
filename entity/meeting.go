package entity

import (
	"encoding/json"
	"io"
	"time"
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
	for _, meeting := range meetings {
		encoder.Encode(meeting)
	}
}

func DeserializeMeeting(r io.Reader) Meetings {
	decoder := json.NewDecoder(r)
	meetings := make(Meetings)
	var meeting Meeting
	for {
		if err := decoder.Decode(&meeting); err == io.EOF {
			return meetings
		}
		meetings.Add(&meeting)
	}
}
