package main

import (
	"crazy_book/src/controllers"
	"fmt"
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
	beego.Router("/updateGrade", &controllers.MainController{}, "post:UpdateGrade")
	beego.Router("/addMyQuestion", &controllers.MainController{}, "post:AddMyQuestion")
	beego.Router("/getMyAllQuestion", &controllers.MainController{}, "get:GetMyAllQuestion")
	beego.Router("/getQuestionById", &controllers.MainController{}, "get:GetQuestionById")
	beego.Router("/getQuestionList", &controllers.MainController{}, "get:GetQuestionList")
	beego.Router("/updateQuestion", &controllers.MainController{}, "post:UpdateQuestion")
	beego.Router("/deletedMyQuestion", &controllers.MainController{}, "post:DeletedMyQuestion")
	beego.Router("/addQuestionComment", &controllers.MainController{}, "post:AddQuestionComment")
	beego.Router("/getQuestionComment", &controllers.MainController{}, "get:GetQuestionComment")

	beego.Router("/addCollection", &controllers.MainController{}, "post:AddCollection")
	beego.Router("/getCollectionQuestionList", &controllers.MainController{}, "get:GetCollectionQuestionList")
	beego.Router("/cancelCollection", &controllers.MainController{}, "post:CancelCollection")
	beego.Router("/addLiked", &controllers.MainController{}, "post:AddLiked")
	beego.Router("/cancelLiked", &controllers.MainController{}, "post:CancelLiked")
	beego.Router("/addLabel", &controllers.MainController{}, "post:AddLabel")
	beego.Router("/deleteLabel", &controllers.MainController{}, "post:DeleteLabel")
	beego.Router("/getUserLabel", &controllers.MainController{}, "get:GetUserLabel")

	//config
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	logs.EnableFuncCallDepth(true)
	logs.SetLevel(logs.LevelDebug)
	logs.SetLogger(logs.AdapterFile, `{"filename":"./logs/book.log"}`)

	// 数据库
	//password := "1234567890"
	password := "123456"
	dataSource := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4", "root", password, "book")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource)
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.DefaultTimeLoc = time.Local

	//listen
	beego.Run("0.0.0.0:8999")
}
