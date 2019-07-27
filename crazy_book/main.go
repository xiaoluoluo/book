package main

import (
	"crazy_book/src/controllers"
	_ "github.com/Go-SQL-Driver/MYSQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

func main() {

	beego.Router("/login", &controllers.MainController{}, "get:Login")
	beego.Router("/getWxOpenId", &controllers.MainController{}, "post:GetWxOpenId")
	beego.Router("/register", &controllers.MainController{}, "post:Register")
	beego.Router("/addMyQuestion", &controllers.MainController{}, "post:AddMyQuestion")
	beego.Router("/getMyAllQuestion", &controllers.MainController{}, "get:GetMyAllQuestion")
	beego.Router("/getQuestionById", &controllers.MainController{}, "get:GetQuestionById")
	beego.Router("/getQuestionList", &controllers.MainController{}, "get:GetQuestionList")
	beego.Router("/updateQuestion", &controllers.MainController{}, "post:UpdateQuestion")
	beego.Router("/deletedMyQuestion", &controllers.MainController{}, "post:DeletedMyQuestion")

	//config
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	logs.EnableFuncCallDepth(true)
	logs.SetLogger(logs.AdapterFile, `{"filename":"./logs/book.log"}`)

	// 数据库
	// online = Sj147258#^(
	// xinzhi = 1234567890
	// luo = 123456

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:1234567890@/book?charset=utf8mb4")
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.DefaultTimeLoc = time.Local

	//listen
	beego.Run("0.0.0.0:8999")
}
