package internal

import "knowtime/internal/chat"

type (
	// BaseMsg 统一的消息响应结构
	BaseMsg struct {
		ErrCode int    `json:"errcode"` // 错误码，0表示成功
		Message string `json:"message"` // 错误消息
	}

	// Response 统一的HTTP响应结构
	Response struct {
		ErrCode int    `json:"errcode"`        // 错误码，0表示成功
		Message string `json:"message"`        // 消息内容
		Data    any    `json:"data,omitempty"` // 响应数据
	}

	InternalUsualMsgPostReq struct {
		AppName  string `json:"app_name"`
		Duration int32  `json:"duration"`
	}
	InternalGenerateReq struct {
		UId uint
		// format yyyy-mm-dd
		Date string `json:"data"`
	}
	InternalChatReq struct {
		UId     uint
		History chat.Messages `json:"history"`
		Message string        `json:"message"`
	}
	InternalGenerateResp struct {
		Output string `json:"output"`
	}
	InternalChatResp struct {
		Output string `json:"output"`
	}
)

// 使用Argon2算法对密码进行哈希处理
type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}
