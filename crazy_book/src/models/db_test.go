package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"testing"
)

func TestInSql(t *testing.T) {
	sql := new(orm.MySQLQueryBuilder).Select("user_id", "user_wid", "user_name", "user_head_pic", "user_grade", "insert_time", "ts").
		From(userTable).Where("user_id").In([]string{"1", "3"}...).String()
	fmt.Println(sql)
}
