package common

import (
	"airad/common/base"
)

// ErrorController definition.
type ErrorController struct {
	base.BaseController
}

func (c *ErrorController) Error404() {
	c.Data["json"] = base.BaseResponse{
		ErrCode: 404,
		ErrMsg:  "Not Found",
	}
	c.ServeJSON()
}
func (c *ErrorController) Error401() {
	c.Data["json"] = base.BaseResponse{
		ErrCode: 401,
		ErrMsg:  "Permission denied",
	}
	c.ServeJSON()
}
func (c *ErrorController) Error403() {
	c.Data["json"] = base.BaseResponse{
		ErrCode: 403,
		ErrMsg:  "Forbidden",
	}
	c.ServeJSON()
}
