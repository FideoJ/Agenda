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
	Sponsor      string
	Participants []string
}

type Meetings map[string]*Meeting

func (meetings Meetings) Has(title string) bool {
	return meetings[title] != nil
}

func (meetings Meetings) Add(meeting *Meeting) {
	meetings[meeting.Title] = meeting
}

func (meetings Meetings) Related(username string) Meetings {
	related := make(Meetings)
	for _, meeting := range meetings {
		isParticipant := false

		for _, participant := range meeting.Participants {
			if participant == username {
				isParticipant = true
			}
		}

		if meeting.Sponsor == username || isParticipant {
			related[meeting.Title] = meeting
		}
	}

	return related
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
	var err error

	for {
		meeting := new(Meeting)
		err = decoder.Decode(meeting)
		if err == io.EOF {
			return meetings
		}
		logger.FatalIf(err)
		meetings.Add(meeting)
	}
}
