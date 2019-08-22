package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Liked struct {
	LikedId    uint64    `json:"liked_id"`
	UserId     uint64    `json:"user_id"`
	QuestionId uint64    `json:"question_id"`
	InsertTime time.Time `json:"insert_time"`
	Ts         time.Time `json:"ts"`
}

const likedTable = "liked"

//增加点赞
func (c *Liked) AddLiked(userId uint64, questionId uint64) (insertId int64, err error) {
	qb := new(orm.MySQLQueryBuilder)
	qb.InsertInto(likedTable, "user_id", "question_id").
		Values("?", "?")
	sql := qb.String()
	o := orm.NewOrm()
	rawSeter, err := o.Raw(sql, userId, questionId).Exec()
	if err != nil {
		return
	}
	return rawSeter.LastInsertId()
}

// 获取点赞
func (c *Liked) GetQuestionLiked(userId, questionId uint64) []Liked {
	var like []Liked
	qb := new(orm.MySQLQueryBuilder)
	qb.Select("liked_id", "user_id", "question_id", "insert_time", "ts").
		From(likedTable).
		Where("question_id = ?").And("user_id = ?")
	sql := qb.String()
	orm.NewOrm().Raw(sql, questionId, userId).QueryRows(&like)
	return like
}

// 获取点赞
func (c *Liked) GetLiked(questionId uint64) []Liked {
	var like []Liked
	qb := new(orm.MySQLQueryBuilder)
	qb.Select("liked_id", "user_id", "question_id", "insert_time", "ts").
		From(likedTable).
		Where("question_id = ?")
	sql := qb.String()
	orm.NewOrm().Raw(sql, questionId).QueryRows(&like)
	return like
}
