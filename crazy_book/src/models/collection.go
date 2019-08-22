package models

import (
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

//增加收藏
func (c *Collection) AddCollection(userId uint64, questionId uint64) (insertId int64, err error) {
	qb := new(orm.MySQLQueryBuilder)
	qb.InsertInto(collectionTable, "user_id", "question_id").
		Values("?", "?")
	sql := qb.String()
	o := orm.NewOrm()
	rawSeter, err := o.Raw(sql, userId, questionId).Exec()
	if err != nil {
		return
	}
	return rawSeter.LastInsertId()
}

// 获取收藏
func (c *Collection) GetQuestionCollection(userId uint64, questionId uint64) []Collection {
	var collection []Collection
	qb := new(orm.MySQLQueryBuilder)
	qb.Select("collection_id", "user_id", "question_id", "insert_time", "ts").
		From(collectionTable).
		Where("user_id = ?").And("question_id = ?")
	sql := qb.String()
	orm.NewOrm().Raw(sql, userId, questionId).QueryRows(&collection)
	return collection
}

// 获取收藏
func (c *Collection) GetCollection(userId uint64) []Collection {
	var collection []Collection
	qb := new(orm.MySQLQueryBuilder)
	qb.Select("collection_id", "user_id", "question_id", "insert_time", "ts").
		From(collectionTable).
		Where("user_id = ?")
	sql := qb.String()
	orm.NewOrm().Raw(sql, userId).QueryRows(&collection)
	return collection
}
