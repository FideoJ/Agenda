package storage

import (
	"os"

	"github.com/FideoJ/Agenda/entity"
	"github.com/FideoJ/Agenda/err"
)

func LoadUsers() entity.Users {
	file, _ := os.OpenFile("./tmp/users.json", os.O_RDONLY|os.O_CREATE, 0644)
	users := entity.DeserializeUser(file)
	return users
}

func StoreUsers(users entity.Users) {
	file, _ := os.OpenFile("./tmp/users.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	users.Serialize(file)
}

func Login(user entity.User) {
	file, _ := os.OpenFile("./tmp/login.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	st, _ := entity.CheckLoginState(file)
	if (st) {
		panic(err.AlreadyLogin)
	} else {
		entity.ChangeLoginState(user.Username, file)
	}
}