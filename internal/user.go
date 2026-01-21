package internal

import (
	"fmt"
	"knowtime/database"
)

func UserLoginInternal(name, password string) (uid uint, b BaseMsg, err error) {
	//获取用户表实例
	userEngine := database.NewUser()
	//先查用户存在性
	userFromDB, err := userEngine.GetByName(name)
	if err != nil {
		return 0, NewBaseMsg(ErrUserNotFound), err
	}
	//比较哈希值
	dbhashedPassword := userFromDB.Password
	match, err := decodeHashString(password, dbhashedPassword)
	if err != nil {
		//哈希解码错误
		return 0, NewBaseMsg(ErrInternalServer), err
	}
	if match == false {
		//密码不匹配
		return 0, NewBaseMsg(ErrPasswordError), fmt.Errorf("password error")
	}
	return userFromDB.UId, NewBaseMsg(SUCCESS), nil
}

func UserLogupInternal(name, password string) (uid uint, b BaseMsg, err error) {
	userEngine := database.NewUser()
	//构造用户结构体
	userToDB := database.User{
		Name:     name,
		Password: hashString(password),
	}
	//创建用户
	uid, err = userEngine.Create(&userToDB)
	if err != nil {
		return 0, NewBaseMsg(ErrCreateUser), err
	}
	return uid, NewBaseMsg(SUCCESS), nil
}
