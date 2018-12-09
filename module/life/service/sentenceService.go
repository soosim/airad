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
	sentences, err := model.ListSentence(vo)
	if nil != err {
		logs.Error(err)
		return baseListResponseVO, err
	}
	baseListResponseVO.Size = vo.Size
	baseListResponseVO.Page = vo.Page
	baseListResponseVO.Total, _ = model.GetSentenceTotal(vo)
	baseListResponseVO.List = sentences
	return baseListResponseVO, nil
}
