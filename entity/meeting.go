package entity

import (
	"encoding/json"
	"io"
	"time"

	"github.com/MarshallW906/Agenda/logger"
)

type Meeting struct {
	Title        string
	StartTime    time.Time
	EndTime      time.Time
	Sponsor      string
	Participants []string
}

func (meeting *Meeting) IsParticipant(username string) bool {
	for _, participant := range meeting.Participants {
		if participant == username {
			return true
		}
	}
	return false
}

func (meeting *Meeting) AddParticipant(username string) {
	meeting.Participants = append(meeting.Participants, username)
}

func (meeting *Meeting) RemoveParticipant(username string) {
	len := len(meeting.Participants)
	for i, participant := range meeting.Participants {
		if participant == username {
			meeting.Participants[i] = meeting.Participants[len-1]
			meeting.Participants = meeting.Participants[:len-1]
			return
		}
	}
}

type Meetings map[string]*Meeting

func (meetings Meetings) Has(title string) bool {
	return meetings[title] != nil
}

func (meetings Meetings) Add(meeting *Meeting) {
	meetings[meeting.Title] = meeting
}

func (meetings Meetings) Remove(meeting *Meeting) {
	delete(meetings, meeting.Title)
}

func (meetings Meetings) Related(username string) Meetings {
	related := make(Meetings)
	for _, meeting := range meetings {
		if meeting.Sponsor == username || meeting.IsParticipant(username) {
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
