package entity

import (
	"encoding/json"
	"io"
)

type User struct {
	Username string
	Password string
	Email    string
	Phone    string
}

type LoginData struct {
	Username string
	State bool
}

type Users map[string]*User

func (users Users) Query(username string) *User {
	return users[username]
}

func (users Users) Add(user *User) {
	users[user.Username] = user
}

func (users Users) Serialize(w io.Writer) {
	encoder := json.NewEncoder(w)
	for _, user := range users {
		encoder.Encode(user)
	}
}

func DeserializeUser(r io.Reader) Users {
	decoder := json.NewDecoder(r)
	users := make(Users)
	var user User
	for {
		if err := decoder.Decode(&user); err == io.EOF {
			return users
		}
		users.Add(&user)
	}
}

func CheckLoginState(r io.Reader) (bool, string) {
	decoder := json.NewDecoder(r)
	var cur LoginData
	decoder.Decode(&cur)
	
	return cur.State, cur.Username
}

func ChangeLoginState(username string, w io.Writer) {
	encoder := json.NewEncoder(w)
	var cur = LoginData{
		Username:username,
		State:true,
	}
	encoder.Encode(cur)
}