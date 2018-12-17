package base

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gopkg.in/go-playground/validator.v9"
	"strings"
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

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// valid input data
func (base *BaseController) ValidInputData(vo interface{}) error {
	validate = validator.New()
	err := validate.Struct(vo)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}

		errRes := *ErrInputData
		errRes.ErrMsg += ":"
		for _, err := range err.(validator.ValidationErrors) {
			errRes.ErrMsg += err.Field() + " " + err.ActualTag() + " " + err.Param() + ","
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
		}
		errRes.ErrMsg = strings.TrimSuffix(errRes.ErrMsg, ",")
		base.Data["json"] = errRes
		base.ServeJSON()
		return errors.New(errRes.ErrMsg)
	}
	return nil
}
