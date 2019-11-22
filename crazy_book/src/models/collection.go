package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type Collection struct {
	CollectionId uint64    `json:"collection_id"`
	UserId       uint64    `json:"user_id"`
	QuestionId   uint64    `json:"question_id"`
	InsertTime   time.Time `json:"insert_time"`
	Ts           time.Time `json:"ts"`
}

const collectionTable = "collection"

var collectionField = []string{"collection_id", "user_id", "question_id", "insert_time", "ts"}

//增加收藏
func (c *Collection) AddCollection(userId uint64, questionId uint64) (insertId int64, err error) {
	qb := new(orm.MySQLQueryBuilder)
	qb.InsertInto(collectionTable, "user_id", "question_id").
		Values("?", "?")
	sql := qb.String()
	o := orm.NewOrm()
	rawSeter, err := o.Raw(sql, userId, questionId).Exec()
	if err != nil {
		logs.Error("AddCollection is err:%v sql:%s", err, sql)
		return
	}
	return rawSeter.LastInsertId()
}

//取消收藏
func (c *Collection) CancelCollection(userId, questionId uint64) error {
	qb := new(orm.MySQLQueryBuilder)
	qb.Delete().From(collectionTable).Where("user_id = ?").And("question_id = ?")
	sql := qb.String()
	o := orm.NewOrm()
	_, err := o.Raw(sql, userId, questionId).Exec()
	return err
}

// 获取收藏
func (c *Collection) GetQuestionCollection(userId uint64, questionId uint64) []Collection {
	var collection []Collection
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(collectionField...).
		From(collectionTable).
		Where("user_id = ?").And("question_id = ?")
	sql := qb.String()
	_, err := orm.NewOrm().Raw(sql, userId, questionId).QueryRows(&collection)
	if err != nil {
		logs.Error("GetQuestionCollection is err:%v sql:%s", err, sql)
	}
	return collection
}

// 获取收藏
func (c *Collection) GetCollection(userId uint64) []Collection {
	var collection []Collection
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(collectionField...).
		From(collectionTable).
		Where("user_id = ?")
	sql := qb.String()
	_, err := orm.NewOrm().Raw(sql, userId).QueryRows(&collection)
	if err != nil {
		logs.Error("GetCollection is err:%v sql:%s", err, sql)
	}
	return collection
}
