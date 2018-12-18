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

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
/*func (vo *ListSentenceVO) Valid(v *validation.Validation) {
	if vo.Page > 1000 || vo.Size > 50 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Page或Size", "Page或Size过大")
	}
}*/

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
