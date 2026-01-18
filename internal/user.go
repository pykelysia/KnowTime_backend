package internal

import (
	"fmt"
	"knowtime/database"
)

func UserLoginInternal(name, password string) (uid uint, b BaseMsg, err error) {
	userEngine := database.NewUser()
	userFromDB, err := userEngine.GetByName(name)
	if err != nil {
		return 0, NewBaseMsg(ErrUserNotFound), err
	}
	dbhashedPassword := userFromDB.Password
	match, err := decodeHashString(password, dbhashedPassword)
	if err != nil {
		return 0, NewBaseMsg(ErrInternalServer), err
	}
	if match == false {
		return 0, NewBaseMsg(ErrPasswordError), fmt.Errorf("password error")
	}
	return userFromDB.UId, NewBaseMsg(SUCCESS), nil
}

func UserLogupInternal(name, password string) (uid uint, b BaseMsg, err error) {
	userEngine := database.NewUser()
	userToDB := database.User{
		Name:     name,
		Password: hashString(password),
	}
	uid, err = userEngine.Create(&userToDB)
	if err != nil {
		return 0, NewBaseMsg(ErrCreateUser), err
	}
	return uid, NewBaseMsg(SUCCESS), nil
}
