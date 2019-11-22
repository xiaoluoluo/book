package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Question struct {
	QuestionId    uint64    `json:"question_id"`
	UserId        uint64    `json:"user_id"`
	UserGrade     uint32    `json:"user_grade"`
	QuestionTitle string    `json:"question_title"`
	AnswerPic     string    `json:"answer_pic"`
	SubjectCode   uint32    `json:"subject_code"`
	TrueTitle     string    `json:"true_title"`
	TruePic       string    `json:"true_pic"`
	FalseTitle    string    `json:"false_title"`
	FalsePic      string    `json:"false_pic"`
	InsertTime    time.Time `json:"insert_time"`
	Ts            time.Time `json:"ts"`
}

const questionTable = "question"

var questionField = []string{"question_id", "user_id", "user_grade", "question_title", "answer_pic", "subject_code", "true_title", "true_pic", "false_title", "false_pic", "insert_time", "ts"}

// 增加我的错题
func (q *Question) AddMyQuestion(userId uint64, userGrade uint32, questionTitle string, answerPic string, subjectCode uint32, trueTitle, truePic, falseTitle, falsePic string) (insertId int64, err error) {
	qb := new(orm.MySQLQueryBuilder)
	qb.InsertInto(questionTable, "user_id", "user_grade", "question_title", "answer_pic", "subject_code", "true_title", "true_pic", "false_title", "false_pic").
		Values("?", "?", "?", "?", "?", "?", "?", "?", "?")
	sql := qb.String()
	o := orm.NewOrm()
	rawSeter, err := o.Raw(sql, userId, userGrade, questionTitle, answerPic, subjectCode, trueTitle, truePic, falseTitle, falsePic).Exec()
	return rawSeter.LastInsertId()
}

//更新题目信息
func (q *Question) UpdateQuestion(questionId, userId uint64, questionTitle, answerPic string, subjectCode uint32, trueTitle, truePic, falseTitle, falsePic string) error {
	qb := new(orm.MySQLQueryBuilder)
	qb.Update(questionTable).
		Set("user_id=?", "question_title = ?", "answer_pic = ?", "subject_code = ?", "true_title = ?", "true_pic = ?", "false_title = ?", "false_pic = ?").Where("question_id = ?")
	sql := qb.String()
	o := orm.NewOrm()
	_, error := o.Raw(sql, userId, questionTitle, answerPic, subjectCode, trueTitle, truePic, falseTitle, falsePic, questionId).Exec()
	return error
}

// 获取我的所有错题
func (q *Question) GetMyAllQuestion(userId uint64, userGrade uint32, limit int, page int) []Question {
	var questions []Question
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(questionField...).
		From(questionTable).
		Where("user_id = ?").And("user_grade = ?")
	qb.Limit(limit).Offset(page)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql, userId, userGrade).QueryRows(&questions)
	return questions
}

//根据科目获取我的错题
func (q *Question) GetMyQuestionBySubject(userId uint64, userGrade uint32, subjectCode uint32, limit int, page int) []Question {
	var questions []Question
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(questionField...).
		From(questionTable).
		Where("user_id = ?").And("user_grade = ?").And("subject_code = ?")
	qb.Limit(limit).Offset(page)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql, userId, userGrade, subjectCode).QueryRows(&questions)
	return questions
}

// 根据题目id 获取题目信息
func (q *Question) GetQuestionById(questionId uint64) []Question {
	var questions []Question
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(questionField...).
		From(questionTable).
		Where("question_id = ?")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql, questionId).QueryRows(&questions)
	return questions
}

//广场中的所有错题
func (q *Question) GetQuestionList(limit int, page int) []Question {
	var questions []Question
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(questionField...).
		From(questionTable)
	qb.Limit(limit).Offset(page)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&questions)
	return questions
}

// 广场中根据年级和科目获取错题
func (q *Question) GetQuestionByGradeAndSubject(userGrade uint32, subjectCode uint32, limit int, page int) []Question {
	var questions []Question
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(questionField...).
		From(questionTable).
		Where("user_grade = ?").And("subject_code = ?")
	qb.Limit(limit).Offset(page)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql, userGrade, subjectCode).QueryRows(&questions)
	return questions
}

// 删除我的题目
func (q *Question) DeletedMyQuestion(questionId uint64) error {
	qb := new(orm.MySQLQueryBuilder)
	qb.Delete().From(questionTable).Where("question_id = ?")
	sql := qb.String()
	o := orm.NewOrm()
	_, err := o.Raw(sql, questionId).Exec()
	return err
}

// 获取题目列表
func (q *Question) GetQuestionListByIds(questionIds []uint64) []Question {
	var questions []Question
	ids := make([]string, 0, len(questionIds))
	for _, id := range questionIds {
		ids = append(ids, strconv.Itoa(int(id)))
	}
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(questionField...).From(questionTable).Where("question_id").In(ids...)
	sql := qb.String()
	orm.NewOrm().Raw(sql).QueryRows(&questions)
	return questions
}
