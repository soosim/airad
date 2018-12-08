package base

// Controller Response is controller Error info struct.
type BaseResponse struct {
	Status  int         `json:"status"`
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

var (
	SuccessReturn   = &BaseResponse{200, 0, "ok", "ok"}
	Err404          = &BaseResponse{404, 404, "找不到网页", "找不到网页"}
	ErrInputData    = &BaseResponse{400, 10001, "数据输入错误", "客户端参数错误"}
	ErrDatabase     = &BaseResponse{500, 10002, "服务器错误", "数据库操作错误"}
	ErrUserToken    = &BaseResponse{500, 10002, "服务器错误", "令牌操作错误"}
	ErrDupUser      = &BaseResponse{400, 10003, "用户信息已存在", "数据库记录重复"}
	ErrNoUser       = &BaseResponse{400, 10004, "用户信息不存在", "数据库记录不存在"}
	ErrPass         = &BaseResponse{400, 10005, "用户信息不存在或密码不正确", "密码不正确"}
	ErrNoUserOrPass = &BaseResponse{400, 10006, "用户不存在或密码不正确", "数据库记录不存在或密码不正确"}
	ErrNoUserChange = &BaseResponse{400, 10007, "用户不存在或数据未改变", "数据库记录不存在或数据未改变"}
	ErrInvalidUser  = &BaseResponse{400, 10008, "用户信息不正确", "Session信息不正确"}
	ErrOpenFile     = &BaseResponse{500, 10009, "服务器错误", "打开文件出错"}
	ErrWriteFile    = &BaseResponse{500, 10010, "服务器错误", "写文件出错"}
	ErrSystem       = &BaseResponse{500, 10011, "服务器错误", "操作系统错误"}
	ErrExpired      = &BaseResponse{400, 10012, "登录已过期", "验证token过期"}
	ErrPermission   = &BaseResponse{400, 10013, "没有权限", "没有操作权限"}
)
