package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
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

func (u *User) Register(userWid string, UserName string, UserHeadPic string) (insertId int64, err error) {
	qb := new(orm.MySQLQueryBuilder)

	qb.InsertInto(userTable, "user_wid", "user_name", "user_head_pic").
		Values("?", "?", "?")
	//返回sql语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	rawSeter, err := o.Raw(sql, userWid, UserName, UserHeadPic).Exec()
	return rawSeter.LastInsertId()
}

func (u *User) Login(userWid string) []User {
	var users []User
	qb := new(orm.MySQLQueryBuilder)

	qb.Select("user_id", "user_wid", "user_name", "user_head_pic", "user_grade", "insert_time", "ts").
		From(userTable).
		Where("user_wid = ?")
	//返回sql语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql, userWid).QueryRows(&users)
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
	qb.Select("user_id", "user_wid", "user_name", "user_head_pic", "user_grade", "insert_time", "ts").
		From(userTable).
		Where("user_id = ?")
	//返回sql语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql, userId).QueryRows(&users)
	return users
}

func (u *User) GetUserList(userIds []uint64) []User {
	var users []User
	ids := make([]string, 0, len(userIds))
	for _, id := range userIds {
		ids = append(ids, strconv.Itoa(int(id)))
	}
	sql := "SELECT user_id, user_wid, user_name, user_head_pic,user_grade, insert_time, ts FROM user WHERE user_id in ("
	sql += strings.Join(ids, ",")
	sql += ")"
	orm.NewOrm().Raw(sql).QueryRows(&users)
	return users
}
