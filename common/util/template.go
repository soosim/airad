package util

import (
	"github.com/astaxie/beego"
)

func InitTemplate() {
	//beego.AddFuncMap("getUsername", models.GetUsername)

	beego.AddFuncMap("getDate", GetDate)
	beego.AddFuncMap("getDateMH", GetDateMH)

	beego.AddFuncMap("getOs", GetOs)
	beego.AddFuncMap("getBrowser", GetBrowser)
	beego.AddFuncMap("getAvatarSource", GetAvatarSource)
	beego.AddFuncMap("getAvatar", GetAvatar)
}
