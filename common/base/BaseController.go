package base

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type BaseController struct {
	beego.Controller
}

func (base *BaseController) Prepare() {
	controllerName, actionName := base.GetControllerAndAction()
	logs.Info("calling :"+controllerName, "/", actionName)
}

func (base *BaseController) URLMapping() {
}

func (base *BaseController) Success(data interface{}) {
	base.Ctx.ResponseWriter.WriteHeader(200)
	base.Data["json"] = BaseResponse{200, 0, "", data}
	base.ServeJSON()
}

// RetError return error information in JSON.
func (base *BaseController) RetError(e *BaseResponse) {
	if mode := beego.AppConfig.String("runmode"); mode == "prod" {
		e.Data = ""
	}

	base.Ctx.ResponseWriter.WriteHeader(e.Status)
	base.Data["json"] = e
	base.ServeJSON()
}

// valid input data
func (base *BaseController) ValidInputData(vo interface{}) {
	if mode := beego.AppConfig.String("runmode"); mode == "prod" {
		// TODO
	}
}