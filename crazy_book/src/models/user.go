package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type User struct {
	UserId      uint64    `json:"user_id"`
	UserWid     string    `json:"user_wid"`
	UserName    string    `json:"user_name"`
	UserHeadPic string    `json:"user_head_pic"`
	UserGrade   uint32    `json:"user_grade"`
	InsertTime  time.Time `json:"insert_time"`
	Ts          time.Time `json:"ts"`
}

const userTable = "user"

var userField = []string{"user_id", "user_wid", "user_name", "user_head_pic", "user_grade", "insert_time", "ts"}

func (u *User) Register(userWid string, UserName string, UserHeadPic string) (insertId int64, err error) {
	qb := new(orm.MySQLQueryBuilder)

	qb.InsertInto(userTable, "user_wid", "user_name", "user_head_pic").
		Values("?", "?", "?")
	sql := qb.String()
	o := orm.NewOrm()
	result, err := o.Raw(sql, userWid, UserName, UserHeadPic).Exec()
	return result.LastInsertId()
}

func (u *User) Login(userWid string) []User {
	var users []User
	qb := new(orm.MySQLQueryBuilder)

	qb.Select(userField...).
		From(userTable).
		Where("user_wid = ?")
	//返回sql语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	_, err := o.Raw(sql, userWid).QueryRows(&users)
	if err != nil {
		logs.Error("Login is err:%v sql:%s", err, sql)
	}
	return users
}

func (u *User) UpdateUserGrade(userId uint64, userGrade uint32) error {

	qb := new(orm.MySQLQueryBuilder)
	qb.Update(userTable).Set("user_grade=?").Where("user_id = ?")
	sql := qb.String()
	o := orm.NewOrm()
	_, error := o.Raw(sql, userGrade, userId).Exec()
	return error
}

func (u *User) GetUserById(userId uint64) []User {
	var users []User
	qb := new(orm.MySQLQueryBuilder)
	qb.Select(userField...).
		From(userTable).
		Where("user_id = ?")
	sql := qb.String()
	_, err := orm.NewOrm().Raw(sql, userId).QueryRows(&users)
	if err != nil {
		logs.Error("GetUserById is err:%v sql:%s", err, sql)
	}
	return users
}

func (u *User) GetUserList(userIds []uint64) []User {
	var users []User
	ids := make([]string, 0, len(userIds))
	for _, id := range userIds {
		ids = append(ids, strconv.Itoa(int(id)))
	}
	qb := new(orm.MySQLQueryBuilder)
	sql := qb.Select(userField...).From(userTable).Where("user_id").In(ids...).String()
	_, err := orm.NewOrm().Raw(sql).QueryRows(&users)
	if err != nil {
		logs.Error("GetUserList is err:%v sql:%s", err, sql)
	}
	return users
}
