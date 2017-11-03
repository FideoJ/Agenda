package entity

import (
	"encoding/json"
	"io"

	"github.com/MarshallW906/Agenda/logger"
)

type User struct {
	Username string
	Password string
	Email    string
	Phone    string
}

type Users map[string]*User

func (users Users) Has(username string) bool {
	return users[username] != nil
}

func (users Users) Query(username string) *User {
	return users[username]
}

func (users Users) Add(user *User) {
	users[user.Username] = user
}

func (users Users) Remove(user *User) {
	delete(users, user.Username)
}

func (users Users) Serialize(w io.Writer) {
	encoder := json.NewEncoder(w)
	var err error

	for _, user := range users {
		err = encoder.Encode(user)
		logger.FatalIf(err)
	}
}

func DeserializeUser(r io.Reader) Users {
	decoder := json.NewDecoder(r)
	users := make(Users)
	var err error

	for {
		user := new(User)
		err = decoder.Decode(user)
		if err == io.EOF {
			return users
		}
		logger.FatalIf(err)
		users.Add(user)
	}
}
