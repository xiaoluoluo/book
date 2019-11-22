package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Question struct {
	QuestionId    uint64    `json:"question_id"`
	UserId        uint64    `json:"user_id"`
	UserGrade     uint32    `json:"user_grade"`
	QuestionTitle string    `json:"question_title"`
	AnswerPic     string    `json:"question_pic"`
	SubjectCode   uint32    `json:"subject_code"`
	TruePic       string    `json:"true_pic1"`
	FalsePic      string    `json:"true_pic2"`
	Point         string    `json:"point"`
	InsertTime    time.Time `json:"insert_time"`
	Ts            time.Time `json:"ts"`
}

const questionTable = "question"

var questionField = []string{"question_id", "user_id", "user_grade", "question_title", "question_pic", "subject_code", "true_pic1", "true_pic2", "point", "insert_time", "ts"}

// 增加我的错题
func (q *Question) AddMyQuestion(userId uint64, userGrade uint32, questionTitle string, questionPic string, subjectCode uint32, truePic1, truePic2, point string) (insertId int64, err error) {
	qb := new(orm.MySQLQueryBuilder)
	qb.InsertInto(questionTable, "user_id", "user_grade", "question_title", "question_pic", "subject_code", "true_pic1", "true_pic2", "point").
		Values("?", "?", "?", "?", "?", "?", "?", "?")
	sql := qb.String()
	o := orm.NewOrm()
	result, err := o.Raw(sql, userId, userGrade, questionTitle, questionPic, subjectCode, truePic1, truePic2, point).Exec()
	return result.LastInsertId()
}

//更新题目信息
func (q *Question) UpdateQuestion(questionId, userId uint64, questionTitle, questionPic string, subjectCode uint32, truePic1, truePic2, point string) error {
	qb := new(orm.MySQLQueryBuilder)
	qb.Update(questionTable).
		Set("user_id=?", "question_title = ?", "question_pic = ?", "subject_code = ?", "true_pic1 = ?", "true_pic2 = ?", "point = ?").Where("question_id = ?")
	sql := qb.String()
	o := orm.NewOrm()
	_, error := o.Raw(sql, userId, questionTitle, questionPic, subjectCode, truePic1, truePic2, point, questionId).Exec()
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
	_, err := o.Raw(sql, userId, userGrade).QueryRows(&questions)
	if err != nil {
		logs.Error("GetMyAllQuestion is err:%v sql:%s", err, sql)
		return questions
	}
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
	_, err := o.Raw(sql, userId, userGrade, subjectCode).QueryRows(&questions)
	if err != nil {
		logs.Error("GetMyQuestionBySubject is err:%v sql:%s", err, sql)
		return questions
	}
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
	_, err := o.Raw(sql, questionId).QueryRows(&questions)
	if err != nil {
		logs.Error("GetQuestionById is err:%v sql:%s", err, sql)
		return questions
	}
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
	_, err := o.Raw(sql).QueryRows(&questions)
	if err != nil {
		logs.Error("GetQuestionList is err:%v sql:%s", err, sql)
		return questions
	}
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
	_, err := o.Raw(sql, userGrade, subjectCode).QueryRows(&questions)
	if err != nil {
		logs.Error("GetQuestionByGradeAndSubject is err:%v sql:%s", err, sql)
		return questions
	}
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
	_, err := orm.NewOrm().Raw(sql).QueryRows(&questions)
	if err != nil {
		logs.Error("GetQuestionListByIds is err:%v sql:%s", err, sql)
		return questions
	}
	return questions
}
