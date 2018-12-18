package controller

import (
	"airad/common/base"
	"airad/common/support"
	"airad/module/life/model"
	"airad/module/life/service"
	sentenceVo "airad/module/life/vo"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

type SentenceController struct {
	base.BaseController
}

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
	if err := c.ValidInputData(vo); err != nil {
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
	if err != nil {
		logs.Error("get sentence error :", err)
		c.RetError(base.ErrServerDatabase)
	} else {
		c.Success(sentence)
	}

	jsonRes, err := json.Marshal(sentence)
	logs.Debug(sentence.Id, jsonRes)
	aaa := &model.Sentence{}
	json.Unmarshal(jsonRes, aaa)
	logs.Debug(aaa)
	if nil != err {
		logs.Error(err)
		return
	}
	_, err = support.GetRedisClient().Set(string(sentence.Id), jsonRes, 0).Result()
	if err != nil {
		logs.Error(err)
	}
	return
}

// @Title Create
// @Description create sentence
// @Success 200 {object} models.Sentence
// @router /create [post]
func (c *SentenceController) Create() {
	logs.Info("接收到的数据为:" + string(c.Ctx.Input.RequestBody))
	vo := sentenceVo.NewSaveSentenceVO()
	if err := base.ParseJsonRequestToVO(c.Ctx, vo); err != nil {
		return
	}
	logs.Info(vo)
	if vo.Id != 0 {
		c.Data["json"] = base.ErrInputData
		return
	}
	sentence, err := service.NewSentenceService().Create(vo)
	if err != nil {
		logs.Error(err)
	}
	c.Success(sentence)
}

// @Title Update
// @Description update sentence
// @Success 200 {object} models.Sentence
// @router /update [post]
func (c *SentenceController) UpdateById() {
	logs.Info("接收到的数据为:", string(c.Ctx.Input.RequestBody))
	vo := sentenceVo.NewSaveSentenceVO()
	if err := base.ParseJsonRequestToVO(c.Ctx, vo); err != nil {
		return
	}
	logs.Info(vo)
	if vo.Id == 0 {
		c.RetError(base.ErrInputData)
		return
	}
	sentence, err := service.NewSentenceService().Update(vo)
	if err != nil {
		logs.Error(err)
	}
	c.Success(sentence)
}
