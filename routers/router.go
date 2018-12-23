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

	demoController "airad/module/demo/controller"
	lifeController "airad/module/life/controller"
)

func init() {
	beego.Router("/", &demoController.MainController{})
	beego.Router("/login", &demoController.UserController{}, "post:Login")
	// beego.Router("/user", &demoControllers.UserController{}, "post:Post")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&demoController.UserController{},
			),
		),
		beego.NSNamespace("/life",
			beego.NSInclude(
				&lifeController.SentenceController{},
				&lifeController.TestController{},
			),
		),
		beego.NSNamespace("/airad",
			beego.NSInclude(
				&demoController.AirAdController{},
			),
		),
		beego.NSNamespace("/device",
			beego.NSInclude(
				&demoController.DeviceController{},
			),
		),
	)
	beego.AddNamespace(ns)
	//beego.InsertFilter("/permission/list", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/v1/device/getdevicebyuserid", &demoController.DeviceController{}, "POST:GetDevicesByUserId")
}
