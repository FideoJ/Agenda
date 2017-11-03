package service

import (
	"fmt"

	"github.com/MarshallW906/Agenda/entity"
	"github.com/MarshallW906/Agenda/err"
	"github.com/MarshallW906/Agenda/logger"
	"github.com/MarshallW906/Agenda/storage"
)

func Register(username string, password string, email string, phone string) {
	users := storage.LoadUsers()

	if users.Has(username) {
		logger.FatalIf(err.UserAlreadyExists)
	}

	users.Add(&entity.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	})
	storage.StoreUsers(users)
}

func Login(username string, password string) {
	Logout()

	users := storage.LoadUsers()
	user := users.Query(username)

	if user == nil || user.Password != password {
		logger.FatalIf(err.WrongUsernameOrPassword)
	}

	storage.StoreSession(&entity.Session{
		CurrentUser: username,
	})
}

func Logout() {
	storage.RemoveSessionFile()
}

func ListAllUsers() {
	_, loggedIn := storage.LoadCurUser()
	if !loggedIn {
		logger.FatalIf(err.RequireLoggedIn)
	}

	users := storage.LoadUsers()

	fmt.Printf("%-20s %-20s %-20s\n", "USERNAME", "EMAIL", "PHONE")
	for _, user := range users {
		fmt.Printf("%-20s %-20s %-20s\n", user.Username, user.Email, user.Phone)
	}
}

func RemoveUser(username string, password string) {
	users := storage.LoadUsers()
	meetings := storage.LoadMeetings()

	user := users.Query(username)

	if user == nil || user.Password != password {
		logger.FatalIf(err.WrongUsernameOrPassword)
	}

	curUser, _ := storage.LoadCurUser()
	if curUser == username {
		Logout()
	}
	users.Remove(user)

	for _, meeting := range meetings {
		if meeting.Sponsor == username {
			meetings.Remove(meeting)
		}
		if meeting.IsParticipant(username) {
			meeting.RemoveParticipant(username)
			if len(meeting.Participants) == 0 {
				meetings.Remove(meeting)
			}
		}
	}

	storage.StoreMeetings(meetings)
	storage.StoreUsers(users)
}
