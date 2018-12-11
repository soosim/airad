package model

import (
	"airad/common/support"
	"airad/module/life/vo"
	"github.com/jinzhu/gorm"
)

type Sentence struct {
	Id      int    `json:"id" gorm:"column(id);pk;unique;auto_increment;int(11)"`
	Content string `json:"content" gorm:"column(content);text"`
	Article string `json:"article" gorm:"column(article);varchar(255)"`
	Role    string `json:"role" gorm:"column(role);varchar(255)"`
	Author  string `json:"author" gorm:"column(author);varchar(20)"`
	Country string `json:"country" gorm:"column(country);varchar(20)"`
	Others  string `json:"others" gorm:"column(others);varchar(1000)"`
	Time    int64  `json:"time" gorm:"column(time);int(11)"`
}

func (u *Sentence) TableName() string {
	return "sentence"
}

func getDBConn() *gorm.DB {
	db, err := support.GetMysqlConnInstance().GetDBConn("life")
	if err != nil {
	}
	return db
}

// 根据用户ID获取用户
func GetSentenceById(id int) (sentence *Sentence, err error) {
	db := getDBConn()
	err = db.First(sentence, id).Error
	return sentence, err
}

func ListSentence(vo *vo.ListSentenceVO) ([]*Sentence, int64, error) {
	var sentences []*Sentence
	db := getDBConn()
	if vo.Content != "" {
		db = db.Where("content LIKE ?", "%"+vo.Content+"%")
	}
	if vo.Author != "" {
		db = db.Where("author LIKE ?", "%"+vo.Author+"%")
	}
	if vo.Country != "" {
		db = db.Where("country = ?", vo.Country)
	}
	var num int64
	err := db.Order("id desc").Offset((vo.Page - 1) * vo.Size).Limit(vo.Size).Find(&sentences).Limit(-1).Count(&num).Error
	return sentences, num, err
}
