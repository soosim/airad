package controller

import (
	"airad/common/base"
	"airad/common/support"
	"airad/module/life/service"
	sentenceVo "airad/module/life/vo"
	"github.com/astaxie/beego/logs"
)

type SentenceController struct {
	base.BaseController
}

// use a single instance of Validate, it caches struct info
// var validate *validator.Validate

// @Title ListSentence
// @Description get all Sentence
// @Success 200 {object} base.BaseListResponseVO
// @router /list [post]
func (c *SentenceController) ListSentence() {
	logs.Info("接收到的数据为:" + string(c.Ctx.Input.RequestBody))
	vo := sentenceVo.NewListSentenceVO()
	if err := base.ParseJsonRequestToVO(c.Ctx, vo); err != nil {
		return
	}
	logs.Debug(vo)
	// err := validator.New().Struct(vo)
	if err := c.ValidInputData(vo); err != nil {
		return
	}
	/*valid := validation.Validation{}
	if ok, err := valid.Valid(vo); nil != err || !ok {
		logs.Debug("has Error")
		if nil != err {
			c.Data["json"] = base.ErrInputData
		} else {
			logs.Error(valid.Errors)
			c.Data["json"] = base.ErrInputData
			for _, err := range valid.Errors {
				c.Data["json"] = base.BaseResponse{400, 400, err.Error(), ""}
				break
			}
		}
		c.ServeJSON()
		return
	}*/
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
	redis := support.GetRedisClient()
	// redis.Set("xiejinlong", "test", 0)
	// redis.Ping()
	redisGet := redis.Get("xiejinlong")
	res, err := redisGet.Result()
	if err != nil {
		logs.Error(err)
	}
	c.Success(res)
	return
	if err == nil {
		c.Success(sentence)
		return
	} else {
		logs.Error("get sentence error :", err)
	}
}

// @Title Create
// @Description create sentence
// @Success 200 {object} models.Sentence
// @router /create [post]
func (c *SentenceController) Create() {
	logs.Debug("接收到的数据为:" + string(c.Ctx.Input.RequestBody))
	vo := sentenceVo.NewSaveSentenceVO()
	if err := base.ParseJsonRequestToVO(c.Ctx, vo); err != nil {
		return
	}
	logs.Debug(vo)
	if vo.Id != 0 {
		c.Data["json"] = base.ErrInputData
	}
	sentence, err := service.NewSentenceService().Create(vo)
	if err == nil {
		c.Success(sentence)
	}
	logs.Error(err)
}
