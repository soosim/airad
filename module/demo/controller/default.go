package controller

import (
	"airad/common/base"
)

// MainController definition.
type MainController struct {
	base.BaseController
}

// Get method.
func (c *MainController) Get() {
	c.Data["Website"] = "www.liudp.cn"
	c.Data["Email"] = "rubinliu@hotmail.com"
	c.TplName = "index.tpl"
	c.Render()
}
