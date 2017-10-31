package service

import (
	"fmt"

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

func Query(username string, password string) error {
	users := storage.LoadUsers()

	if username == "" {
		fmt.Println("Current Registered Users:")
		for name := range users {
			fmt.Println(name)
		}
	} else {
		if qusr := users.Query(username); qusr != nil && qusr.Password == password {
			fmt.Printf("username: %+v, e-mail: %+v, phone: %+v\n", qusr.Username, qusr.Email, qusr.Phone)
		} else {
			fmt.Print("Cannot find this user. check the username and password.")
		}
	}

	return nil
}
