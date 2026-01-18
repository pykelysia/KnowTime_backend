package internal

type (
	BaseMsg struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
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
	InternalGenerateResp struct {
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
