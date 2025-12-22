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
		// format yyyy-mm-dd
		Date string `json:"data"`
	}
	InternalGenerateResp struct {
		Output string `json:"output"`
	}
)
