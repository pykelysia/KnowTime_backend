package route

import "knowtime/internal"

type BaseMsg internal.BaseMsg

// UserLoginReq 用户登录请求结构
type UserLoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserLogupReq 用户注册请求结构
type UserLogupReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// LoginData 登录成功响应数据
type LoginData struct {
	Token string `json:"token"`
	UId   uint   `json:"u_id"`
}

// LogupData 注册成功响应数据
type LogupData struct {
	UId uint `json:"u_id"`
}

// UserInfoResp 用户信息响应
type UserInfoResp struct {
	Name string `json:"name"`
	UId  uint   `json:"u_id"`
}
