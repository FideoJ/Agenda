package service

import (
	"time"

	"../entity"
	"../err"
	"../logger"
	"../storage"
	"../utils"
)

func CreateMeeting(title string, startTimeStr string, endTimeStr string, participants []string) {
	curUser, loggedIn := storage.LoadCurUser()
	if !loggedIn {
		logger.FatalIf(err.RequireLoggedIn)
	}

	meetings := storage.LoadMeetings()
	if meetings.Has(title) {
		logger.FatalIf(err.TitleAlreadyExists)
	}

	sTime, timeerr := time.Parse(utils.TimeLayout, startTimeStr)
	logger.FatalIf(timeerr)
	eTime, timeerr := time.Parse(utils.TimeLayout, endTimeStr)
	logger.FatalIf(timeerr)
	if !sTime.Before(eTime) {
		logger.FatalIf(err.StartTimeNotBeforeEndTime)
	}

	users := storage.LoadUsers()
	referredUsers := make(map[string]bool)

	attendants := append([]string(nil), curUser)
	attendants = append(attendants, participants...)

	for _, attendant := range attendants {
		if !users.Has(attendant) {
			logger.FatalIf(err.UserNotExist)
		}
		if _, referred := referredUsers[attendant]; referred {
			logger.FatalIf(err.AttendantsDuplicated)
		}
		referredUsers[attendant] = true

		relatedMeetings := meetings.Related(attendant)
		for _, relatedMeeting := range relatedMeetings {
			if utils.Overlapped(relatedMeeting.StartTime, relatedMeeting.EndTime, sTime, eTime) {
				logger.FatalIf(err.TimeConflicted(attendant, relatedMeeting))
			}
		}
	}

	meetings.Add(&entity.Meeting{
		Title:        title,
		StartTime:    sTime,
		EndTime:      eTime,
		Sponsor:      curUser,
		Participants: participants,
	})

	storage.StoreMeetings(meetings)
}

func CancelMeeting(title string) {
	curUser, loggedIn := storage.LoadCurUser()
	if !loggedIn {
		logger.FatalIf(err.RequireLoggedIn)
	}

	meetings := storage.LoadMeetings()
	for _, meeting := range meetings {
		if meeting.Title == title && meeting.Sponsor == curUser {
			meetings.Remove(meeting)
			return
		}
	}
	logger.FatalIf(err.MeetingNotFound)
}
