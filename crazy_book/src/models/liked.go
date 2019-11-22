package models

import (
	"github.com/astaxie/beego/logs"
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

var likedField = []string{"liked_id", "user_id", "question_id", "insert_time", "ts"}

//增加点赞
func (c *Liked) AddLiked(userId uint64, questionId uint64) (insertId int64, err error) {
	qb := new(orm.MySQLQueryBuilder)
	qb.InsertInto(likedTable, "user_id", "question_id").
		Values("?", "?")
	sql := qb.String()
	o := orm.NewOrm()
	result, err := o.Raw(sql, userId, questionId).Exec()
	if err != nil {
		logs.Error("AddLiked is err:%v sql:%s", err, sql)
		return
	}
	return result.LastInsertId()
}

// 获取点赞
func (c *Liked) GetQuestionLiked(userId, questionId uint64) []Liked {
	var like []Liked
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(likedField...).
		From(likedTable).
		Where("question_id = ?").And("user_id = ?")
	sql := qb.String()
	_, err := orm.NewOrm().Raw(sql, questionId, userId).QueryRows(&like)
	if err != nil {
		logs.Error("GetQuestionLiked is err:%v sql:%s", err, sql)
		return like
	}
	return like
}

// 取消点赞
func (c *Liked) CancelLiked(userId, questionId uint64) error {
	qb := new(orm.MySQLQueryBuilder)
	qb.Delete().From(likedTable).Where("user_id = ?").And("question_id = ?")
	sql := qb.String()
	o := orm.NewOrm()
	_, err := o.Raw(sql, userId, questionId).Exec()
	return err
}

// 获取点赞
func (c *Liked) GetLiked(questionId uint64) []Liked {
	var like []Liked
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(likedField...).
		From(likedTable).
		Where("question_id = ?")
	sql := qb.String()
	_, err := orm.NewOrm().Raw(sql, questionId).QueryRows(&like)
	if err != nil {
		logs.Error("GetLiked is err:%v sql:%s", err, sql)
		return like
	}
	return like
}
