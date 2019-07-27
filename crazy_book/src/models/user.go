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
	InsertTime  time.Time `json:"insert_time"`
	Ts          time.Time `json:"ts"`
}

func (u *User) Register(userWid string, UserName string, UserHeadPic string) (insertId int64, err error) {
	qb := new(orm.MySQLQueryBuilder)

	qb.InsertInto("user", "user_wid", "user_name", "user_head_pic").
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

	qb.Select("user_id", "user_wid", "user_name", "user_head_pic", "insert_time", "ts").
		From("user").
		Where("user_wid = ?")
	//返回sql语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql, userWid).QueryRows(&users)
	return users
}

func (u *User) GetUserList(userIds []uint64) []User {
	var users []User
	ids := make([]string,0,len(userIds))
	for _,id := range  userIds {
		ids =append(ids,strconv.Itoa(int(id)))
	}
	sql := "SELECT user_id, user_wid, user_name, user_head_pic, insert_time, ts FROM user WHERE user_id in ("
	sql += strings.Join(ids, ",")
	sql += ")"
	orm.NewOrm().Raw(sql).QueryRows(&users)
	return users
}