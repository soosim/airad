package base

// Controller Response is controller Error info struct.
type BaseResponse struct {
	Status  int         `json:"status"`
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

var (
	SuccessReturn = &BaseResponse{200, 0, "ok", "ok"}
	Err404        = &BaseResponse{404, 404, "找不到网页", "找不到网页"}

	// 客户端错误
	ErrInputData    = &BaseResponse{400, 10001, "数据输入错误", "客户端参数错误"}
	ErrDupUser      = &BaseResponse{400, 10003, "数据已存在", "数据记录重复"}
	ErrNoUser       = &BaseResponse{400, 10004, "数据不存在", "数据记录不存在"}
	ErrPass         = &BaseResponse{400, 10005, "密码不正确", "密码不正确"}
	ErrNoUserOrPass = &BaseResponse{400, 10006, "用户不存在或密码不正确", "记录不存在或密码不正确"}
	ErrNoUserChange = &BaseResponse{400, 10007, "用户不存在或数据未改变", "记录不存在或数据未改变"}
	ErrInvalidUser  = &BaseResponse{400, 10008, "用户信息不正确", "Session信息不正确"}
	ErrExpired      = &BaseResponse{400, 10012, "登录已过期", "验证token过期"}
	ErrPermission   = &BaseResponse{400, 10013, "没有权限", "没有操作权限"}

	//服务端错误
	ErrServerUnKnown   = &BaseResponse{500, 20001, "服务器错误", "未知错误"}
	ErrServerDatabase  = &BaseResponse{500, 20002, "服务器错误", "数据库操作错误"}
	ErrServerUserToken = &BaseResponse{500, 20003, "服务器错误", "令牌操作错误"}
	ErrServerOpenFile  = &BaseResponse{500, 20009, "服务器错误", "打开文件出错"}
	ErrServerWriteFile = &BaseResponse{500, 20010, "服务器错误", "写文件出错"}
	ErrServerSystem    = &BaseResponse{500, 20011, "服务器错误", "操作系统错误"}
)
