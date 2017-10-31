package service

import (
	"fmt"

	"github.com/MarshallW906/Agenda/entity"
	"github.com/MarshallW906/Agenda/err"
	"github.com/MarshallW906/Agenda/storage"
)

func Register(username string, password string, email string, phone string) error {
	users := storage.LoadUsers()
	if username == "" {
		return err.RegWithEmptyUsername
	}
	if password == "" {
		return err.RegWithEmptyPassword
	}
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

func Delete(username string, password string) error {
	users := storage.LoadUsers()
	if qusr := users.Query(username); qusr != nil && qusr.Password == password {
		users.Delete(qusr)
		// need change to log
		fmt.Printf("Delete User [%+v] Successfully.\n", username)
	} else {
		fmt.Println("Deleted Failed. Please check the username and Password.")
	}
	storage.StoreUsers(users)
	return nil
}
