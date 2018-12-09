package vo

import (
	"airad/common/base"
	"github.com/astaxie/beego/validation"
)

type ListSentenceVO struct {
	base.BaseListRequestVO
	Content string `valid:"MaxSize(10)"`
	Article string
	Role    string
	Author  string
	Country string
	Others  string
	Time    int64
}

func NewListSentenceVO() *ListSentenceVO {
	return &ListSentenceVO{}
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (vo *ListSentenceVO) Valid(v *validation.Validation) {
	if vo.Page > 1000 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Page", "页数过大")
	}
}