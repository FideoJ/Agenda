package service

import (
	"time"

	"github.com/MarshallW906/Agenda/entity"
	"github.com/MarshallW906/Agenda/err"
	"github.com/MarshallW906/Agenda/storage"
)

var (
	timeLayoutStr string
)

func init() {
	timeLayoutStr = "2006-01-02/15:04:05"
}

func CreateMt(title string, startTimeStr string, endTimeStr string, participants []string) error {
	// users := storage.LoadUsers()
	meetings := storage.LoadMeetings()
	if meetings.Query(title) != nil {
		return err.MeetingTitleAlreadyExists
	}

	sTime, terr := time.Parse(timeLayoutStr, startTimeStr)
	if terr != nil {
		return terr
	}
	eTime, terr := time.Parse(timeLayoutStr, endTimeStr)
	if terr != nil {
		return terr
	}

	var sponsorUsr *entity.User

	meetings.Add(&entity.Meeting{
		Title:        title,
		StartTime:    sTime,
		EndTime:      eTime,
		Sponsor:      sponsorUsr,
		Participants: participants,
	})
	storage.StoreMeetings(meetings)
	return nil
}
