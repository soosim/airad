package service

import (
	"airad/common/base"
	"airad/common/helper"
	"airad/module/life/model"
	"airad/module/life/vo"
	"github.com/astaxie/beego/logs"
)

type sentenceService struct {
	base.BaseService
}

func NewSentenceService() *sentenceService {
	return &sentenceService{}
}

func (s *sentenceService) ListSentence(vo *vo.ListSentenceVO) (base.BaseListResponseVO, error) {
	var baseListResponseVO = base.BaseListResponseVO{}
	var err error
	baseListResponseVO.List, baseListResponseVO.Total, err = model.ListSentence(vo)
	if nil != err {
		logs.Error(err)
		return baseListResponseVO, err
	}
	baseListResponseVO.Size = vo.Size
	baseListResponseVO.Page = vo.Page
	return baseListResponseVO, nil
}

func (s *sentenceService) GetOneByRand() (model.Sentence, error) {
	return model.GetOneByRand()
}

func (s *sentenceService) Create(vo *vo.SaveSentenceVO) (model.Sentence, error) {
	sentence := &model.Sentence{}
	if err := helper.Copy(sentence, *vo); err != nil {
		logs.Error("copy from vo to sentence error")
		logs.Error(err)
		return model.Sentence{}, err
	}
	sentenceModel, err := model.Create(*sentence)
	if err == nil {
		return sentenceModel, nil
	}
	logs.Error("保存失败")
	return model.Sentence{}, err
}
