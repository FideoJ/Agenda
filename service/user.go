package service

import (
	"github.com/FideoJ/Agenda/entity"
	"github.com/FideoJ/Agenda/err"
	"github.com/FideoJ/Agenda/storage"
)

func Register(username string, password string, email string, phone string) error {
	users := storage.LoadUsers()
	if users.Query(username) != nil {
		return err.UserAlreadyExists
	}
	users.Add(&entity.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	})
	storage.StoreUsers(users)
	return nil
}

func Query(username string, email string) error {
	// users := storage.LoadUsers()
	return nil
}
