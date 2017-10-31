package service

import (
	"github.com/MarshallW906/Agenda/entity"
	"github.com/MarshallW906/Agenda/err"
	"github.com/MarshallW906/Agenda/storage"
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
