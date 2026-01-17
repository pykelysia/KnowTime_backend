package internal

import (
	"fmt"
	"knowtime/database"
)

func UserLoginInternal(name, password string) (uid uint, b BaseMsg, err error) {
	userEngine := database.NewUser()
	userFromDB, err := userEngine.GetByName(name)
	if err != nil {
		return 0, BaseMsg{500, "Failed to found user"}, err
	}
	match, err := decodeHashString(userFromDB.Password, password)
	if err != nil {
		return 0, BaseMsg{500, "DB Mal-format"}, err
	}
	if match == false {
		return 0, BaseMsg{400, "Password error"}, fmt.Errorf("Password Error")
	}
	return userFromDB.UId, BaseMsg{}, nil
}

func UserLogupInternal(name, password string) (uid uint, b BaseMsg, err error) {
	userEngine := database.NewUser()
	userToDB := database.User{
		Name:     name,
		Password: hashString(password),
	}
	uid, err = userEngine.Create(&userToDB)
	if err != nil {
		return 0, BaseMsg{500, "Could not log up user"}, err
	}
	return uid, BaseMsg{}, nil
}
