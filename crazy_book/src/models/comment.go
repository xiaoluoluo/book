package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Comment struct {
	CommentId    uint64    `json:"comment_id"`
	UserId       uint64    `json:"user_id"`
	QuestionId   uint64    `json:"question_id"`
	CommentIntro string    `json:"comment_intro"`
	InsertTime   time.Time `json:"insert_time"`
	Ts           time.Time `json:"ts"`
}

const commentTable = "comment"

var commentField = []string{"comment_id", "user_id", "question_id", "comment_intro", "insert_time", "ts"}

//增加评论
func (c *Comment) AddComment(userId uint64, questionId uint64, commentIntro string) (insertId int64, err error) {
	qb := new(orm.MySQLQueryBuilder)
	qb.InsertInto(commentTable, "user_id", "question_id", "comment_intro").
		Values("?", "?", "?")
	sql := qb.String()
	o := orm.NewOrm()
	rawSeter, err := o.Raw(sql, userId, questionId, commentIntro).Exec()
	if err != nil {
		return
	}
	return rawSeter.LastInsertId()
}

// 获取评论
func (c *Comment) GetComment(questionId uint64) []Comment {
	var comment []Comment
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(commentField...).
		From(commentTable).
		Where("question_id = ?")
	sql := qb.String()
	orm.NewOrm().Raw(sql, questionId).QueryRows(&comment)
	return comment

}
