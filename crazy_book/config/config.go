package config

import (
	conf "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

var (
	AppId     string
	Secret    string
	Password  string
	DataBases string
	DataUser  string
)

func Init() {
	iniConf, err := conf.NewConfig("ini", "./config/config.conf")
	if err != nil {
		logs.Error("NewConfig is err:%v", err)
		panic("config is err")
	}
	AppId = iniConf.String("wei_xin::AppId")
	Secret = iniConf.String("wei_xin::Secret")
	Password = iniConf.String("db::Password")
	DataBases = iniConf.String("db::DataBases")
	DataUser = iniConf.String("db::DataUser")
}
