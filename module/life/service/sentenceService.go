package service

import (
	"airad/common/base"
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
