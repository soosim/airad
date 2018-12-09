package controller

import (
	"airad/common/base"
	"airad/module/life/service"
	sentenceVo "airad/module/life/vo"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

type SentenceController struct {
	base.BaseController
}

// @Title GetAll
// @Description get all Sentence
// @Success 200 {object} models.Sentence
// @router /list [post]
func (c *SentenceController) ListSentence() {
	logs.Debug("接收到的数据为:" + string(c.Ctx.Input.RequestBody))
	vo := sentenceVo.NewListSentenceVO()
	if err := base.ParseJsonRequestToVO(c.Ctx, vo); err != nil {
		return
	}
	valid := validation.Validation{}
	if hasError, err := valid.Valid(vo); nil != err || hasError {
		if nil != err {
			c.Data["json"] = base.ErrInputData
		} else {
			for _, err := range valid.Errors {
				c.Data["json"] = base.BaseResponse{400, 400, err.Error(), ""}
				break
			}
		}
		c.ServeJSON()
		return
	}
	sentenceResponseVO, err := service.NewSentenceService().ListSentence(vo)
	if nil != err {
		c.Data["json"] = base.ErrDatabase
		c.ServeJSON()
		return
	}
	c.Success(sentenceResponseVO)
}
