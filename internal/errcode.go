package internal

// 错误码定义
const (
	// 成功
	SUCCESS = 0

	// 客户端错误 (1xxx)
	ErrInvalidRequest     = 1001 // 无效的请求参数
	ErrInvalidRequestBody = 1002 // 无效的请求体
	ErrUnauthorized       = 1003 // 未授权
	ErrUserNotFound       = 1004 // 用户不存在
	ErrPasswordError      = 1005 // 密码错误

	// 服务器错误 (2xxx)
	ErrInternalServer     = 2001 // 服务器内部错误
	ErrDatabaseError      = 2002 // 数据库错误
	ErrCreateTimeEvent    = 2003 // 创建时间事件失败
	ErrUpdateTimeEvent    = 2004 // 更新时间事件失败
	ErrGenerateToken      = 2005 // 生成Token失败
	ErrBuildAgent         = 2006 // 构建Agent失败
	ErrCallAgent          = 2007 // 调用Agent失败
	ErrCreateUser         = 2008 // 创建用户失败
	ErrDatabaseConnection = 2009 // 数据库连接失败
)

// 错误消息映射
var ErrCodeMsg = map[int]string{
	SUCCESS: "Success",

	// 客户端错误
	ErrInvalidRequest:     "Invalid request",
	ErrInvalidRequestBody: "Invalid request body",
	ErrUnauthorized:       "User not authenticated",
	ErrUserNotFound:       "User not found",
	ErrPasswordError:      "Password error",

	// 服务器错误
	ErrInternalServer:     "Internal server error",
	ErrDatabaseError:      "Database error",
	ErrCreateTimeEvent:    "Failed to create time event",
	ErrUpdateTimeEvent:    "Failed to update time event",
	ErrGenerateToken:      "Failed to generate token",
	ErrBuildAgent:         "Failed to build ReAct agent",
	ErrCallAgent:          "Failed to call Agent",
	ErrCreateUser:         "Failed to create user",
	ErrDatabaseConnection: "Database connection failed",
}

// GetErrMsg 获取错误消息
func GetErrMsg(code int) string {
	if msg, ok := ErrCodeMsg[code]; ok {
		return msg
	}
	return ErrCodeMsg[ErrInternalServer]
}

// NewResponse 创建统一响应
func NewResponse(errCode int, data any) Response {
	return Response{
		ErrCode: errCode,
		Message: GetErrMsg(errCode),
		Data:    data,
	}
}

// NewBaseMsg 创建基础消息
func NewBaseMsg(errCode int) BaseMsg {
	return BaseMsg{
		ErrCode: errCode,
		Message: GetErrMsg(errCode),
	}
}
