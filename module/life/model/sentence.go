package model

import (
	"airad/module/life/vo"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Sentence struct {
	Id      int    `json:"id" orm:"column(id);pk;unique;auto_increment;int(11)"`
	Content string `json:"content" orm:"column(content);text"`
	Article string `json:"article" orm:"column(article);varchar(255)"`
	Role    string `json:"role" orm:"column(role);varchar(255)"`
	Author  string `json:"author" orm:"column(author);varchar(20)"`
	Country string `json:"country" orm:"column(country);varchar(20)"`
	Others  string `json:"others" orm:"column(others);varchar(1000)"`
	Time    int64  `json:"time" orm:"column(time);varchar(11)"`
}

func (u *Sentence) TableName() string {
	return "sentence"
}

func sentenceQuerySeter() orm.QuerySeter {
	o := orm.NewOrm()
	o.Using("life")
	return o.QueryTable(new(Sentence))
}

// 构建列表查询器
func buildListQuerySeter() {

}

// 根据用户ID获取用户
func GetSentenceById(id int) (v *Sentence, err error) {
	v = &Sentence{Id: id}
	if err = sentenceQuerySeter().Filter("id", id).One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func ListSentence(vo *vo.ListSentenceVO) ([]*Sentence, error) {
	var sentences []*Sentence
	qs := sentenceQuerySeter()
	if vo.Content != "" {
		qs = qs.Filter("content__icontains", vo.Content)
	}
	if vo.Author != "" {
		qs = qs.Filter("author__icontains", vo.Author)
	}
	if vo.Country != "" {
		qs = qs.Filter("country", vo.Country)
	}
	qs.OrderBy("-id")
	num, err := qs.Limit(vo.Size).Offset((vo.Page - 1) * vo.Size).All(&sentences)
	logs.Debug("Returned Rows Num: %d, %s", num, err)
	return sentences, err
}

func GetSentenceTotal(vo *vo.ListSentenceVO) (int64, error) {
	qs := sentenceQuerySeter()
	qs = qs.Filter("content__icontains", vo.Content)
	return qs.Count()
}
