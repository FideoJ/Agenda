package service

import (
	"fmt"

	"../entity"
	"../err"
	"../logger"
	"../storage"
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
	user := users.Query(username)

	if user == nil || user.Password != password {
		logger.FatalIf(err.WrongUsernameOrPassword)
	}

	curUser, _ := storage.LoadCurUser()
	if curUser == username {
		Logout()
	}
	users.Remove(user)
	storage.StoreUsers(users)
}
