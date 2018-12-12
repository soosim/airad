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
	logs.Debug(vo)

	valid := validation.Validation{}
	if ok, err := valid.Valid(vo); nil != err || !ok {
		logs.Debug("has Error")
		if nil != err {
			c.Data["json"] = base.ErrInputData
		} else {
			logs.Error(valid.Errors)
			c.Data["json"] = base.BaseResponse{400, 400, "参数错误", ""}
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
		c.Data["json"] = base.ErrServerDatabase
		c.ServeJSON()
		return
	}
	c.Success(sentenceResponseVO)
}

// @Title GetOneByRand
// @Description get one Sentence by rand
// @Success 200 {object} models.Sentence
// @router /getOneByRand [get]
func (c *SentenceController) GetOneByRand() {
	sentence, err := service.NewSentenceService().GetOneByRand()
	if err == nil {
		c.Success(sentence)
		return
	} else {

	}
}
