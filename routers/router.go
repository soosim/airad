// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"

	"airad/module/demo/controller"
)

func init() {
	beego.Router("/", &controller.MainController{})
	beego.Router("/login", &controller.UserController{}, "post:Login")
	// beego.Router("/user", &controllers.UserController{}, "post:Post")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controller.UserController{},
			),
		),
		beego.NSNamespace("/airad",
			beego.NSInclude(
				&controller.AirAdController{},
			),
		),
		beego.NSNamespace("/device",
			beego.NSInclude(
				&controller.DeviceController{},
			),
		),
	)
	beego.AddNamespace(ns)
	//beego.InsertFilter("/permission/list", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/v1/device/getdevicebyuserid", &controller.DeviceController{}, "POST:GetDevicesByUserId")
}
