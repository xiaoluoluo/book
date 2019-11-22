package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type Label struct {
	LabelId     uint64    `json:"label_id"`
	UserId      uint64    `json:"user_id"`
	SubjectCode uint32    `json:"subject_code"`
	Label       string    `json:"label"`
	InsertTime  time.Time `json:"insert_time"`
	Ts          time.Time `json:"ts"`
}

const labelTable = "label"

var labelField = []string{"label_id", "user_id", "subject_code", "label", "insert_time", "ts"}

//增加标签
func (c *Label) AddUserLabel(userId uint64, subjectCode uint32, label string) (insertId int64, err error) {
	qb := new(orm.MySQLQueryBuilder)
	qb.InsertInto(labelTable, "user_id", "subject_code", "label").
		Values("?", "?", "?")
	sql := qb.String()
	o := orm.NewOrm()
	rawSeter, err := o.Raw(sql, userId, subjectCode, label).Exec()
	if err != nil {
		logs.Error("AddUserLabel is err:%v sql:%s", err, sql)
		return
	}
	return rawSeter.LastInsertId()
}

// 获取标签
func (c *Label) GetUserSubjectLabel(userId uint64, subjectCode uint32) []Label {
	var label []Label
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(labelField...).
		From(labelTable).
		Where("user_id = ?").And("subject_code = ?")
	sql := qb.String()
	_, err := orm.NewOrm().Raw(sql, userId, subjectCode).QueryRows(&label)
	if err != nil {
		logs.Error("GetUserSubjectLabel is err:%v sql:%s", err, sql)
		return label
	}
	return label
}

// 获取标签
func (c *Label) GetUserLabel(userId uint64) []Label {
	var label []Label
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(labelField...).
		From(labelTable).
		Where("user_id = ?")
	sql := qb.String()
	_, err := orm.NewOrm().Raw(sql, userId).QueryRows(&label)
	if err != nil {
		logs.Error("GetUserLabel is err:%v sql:%s", err, sql)
		return label
	}
	return label
}

// 删除我的标签知识点
func (q *Label) DeletedUserLabel(userId, labelId uint64) error {
	qb := new(orm.MySQLQueryBuilder)
	qb.Delete().From(labelTable).Where("label_id = ?").And("user_id = ?")
	sql := qb.String()
	o := orm.NewOrm()
	_, err := o.Raw(sql, labelId, userId).Exec()
	return err
}
