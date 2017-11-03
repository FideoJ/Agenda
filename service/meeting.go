package service

import (
	"fmt"
	"time"

	"github.com/MarshallW906/Agenda/entity"
	"github.com/MarshallW906/Agenda/err"
	"github.com/MarshallW906/Agenda/logger"
	"github.com/MarshallW906/Agenda/storage"
	"github.com/MarshallW906/Agenda/utils"
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
			storage.StoreMeetings(meetings)
			return
		}
	}
	logger.FatalIf(err.MeetingNotFound)
}

func QuitMeeting(title string) {
	curUser, loggedIn := storage.LoadCurUser()
	if !loggedIn {
		logger.FatalIf(err.RequireLoggedIn)
	}

	meetings := storage.LoadMeetings()
	for _, meeting := range meetings {
		if meeting.Title == title && meeting.IsParticipant(curUser) {
			meeting.RemoveParticipant(curUser)
			if len(meeting.Participants) == 0 {
				meetings.Remove(meeting)
			}
			storage.StoreMeetings(meetings)
			return
		}
	}
	logger.FatalIf(err.MeetingNotFound)
}

func ClearMeetings() {
	curUser, loggedIn := storage.LoadCurUser()
	if !loggedIn {
		logger.FatalIf(err.RequireLoggedIn)
	}

	meetings := storage.LoadMeetings()
	for _, meeting := range meetings {
		if meeting.Sponsor == curUser {
			meetings.Remove(meeting)
		}
	}
	storage.StoreMeetings(meetings)
}

func AddParticipant(title string, participants []string) {
	curUser, loggedIn := storage.LoadCurUser()
	if !loggedIn {
		logger.FatalIf(err.RequireLoggedIn)
	}

	users := storage.LoadUsers()
	meetings := storage.LoadMeetings()

	for _, meeting := range meetings {
		if meeting.Title == title && meeting.Sponsor == curUser {
			for _, singleParticipant := range participants {
				if meeting.Sponsor == singleParticipant {
					logger.FatalIf(err.AttendantsDuplicated)
				}

				if users.Has(singleParticipant) {
					if !meeting.IsParticipant(singleParticipant) {
						meeting.AddParticipant(singleParticipant)
					} else {
						logger.FatalIf(err.AttendantsDuplicated)
					}
				} else {
					logger.FatalIf(err.UserNotExist)
				}
			}
			storage.StoreMeetings(meetings)
			return
		}
	}
	logger.FatalIf(err.MeetingNotFound)
}

func RemoveParticipant(title string, participants []string) {
	curUser, loggedIn := storage.LoadCurUser()
	if !loggedIn {
		logger.FatalIf(err.RequireLoggedIn)
	}

	users := storage.LoadUsers()
	meetings := storage.LoadMeetings()

	for _, meeting := range meetings {
		if meeting.Title == title && meeting.Sponsor == curUser {
			for _, singleParticipant := range participants {
				if meeting.Sponsor == singleParticipant {
					logger.FatalIf(err.SponsorRemoveSelf)
				}

				if users.Has(singleParticipant) && meeting.IsParticipant(singleParticipant) {
					meeting.RemoveParticipant(singleParticipant)
				} else {
					logger.FatalIf(err.UserNotExist)
				}
			}
			if len(meeting.Participants) == 0 {
				meetings.Remove(meeting)
			}
			storage.StoreMeetings(meetings)
			return
		}
	}
	logger.FatalIf(err.MeetingNotFound)
}

func ListAllMeetings() {
	curUser, loggedIn := storage.LoadCurUser()
	if !loggedIn {
		logger.FatalIf(err.RequireLoggedIn)
	}

	meetings := storage.LoadMeetings()
	relatedMeetings := meetings.Related(curUser)

	fmt.Printf("%-15s %-15s %-20s %-20s %-12s\n", "TITLE", "SPONSOR", "START-TIME", "END-TIME", "PARTICIPANTS")
	for _, meeting := range relatedMeetings {
		fmt.Printf("%-15s %-15s %-20s %-20s",
			meeting.Title, meeting.Sponsor,
			meeting.StartTime.Format(utils.TimeLayout), meeting.EndTime.Format(utils.TimeLayout))
		for _, singleParticipant := range meeting.Participants {
			fmt.Printf(" %-12s  ", singleParticipant)
		}
		fmt.Printf("\n")
	}
}
