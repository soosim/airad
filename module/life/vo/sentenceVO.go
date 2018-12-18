package vo

import (
	"airad/common/base"
)

type ListSentenceVO struct {
	base.BaseListRequestVO
	Content string `validate:"omitempty,min=1,max=30"`
	Article string `validate:"omitempty,min=1,max=20"`
	Role    string `validate:"omitempty,min=1,max=20"`
	Author  string `validate:"omitempty,min=1,max=20"`
	Country string `validate:"omitempty,min=1,max=20"`
	Others  string
	Time    int64
}

func NewListSentenceVO() *ListSentenceVO {
	return &ListSentenceVO{}
}

type SaveSentenceVO struct {
	Id      int    `json:"id" validate:"omitempty,min=1"`
	Content string `json:"content" validate:"omitempty,min=1,max=30"`
	Article string `json:"article" validate:"omitempty,min=1,max=20"`
	Role    string `json:"role" validate:"omitempty,min=1,max=20"`
	Author  string `json:"author" validate:"omitempty,min=1,max=20"`
	Country string `json:"country" validate:"omitempty,min=1,max=20"`
	Others  string `json:"others"`
	Time    int64  `json:"time"`
}

func NewSaveSentenceVO() *SaveSentenceVO {
	return &SaveSentenceVO{}
}
