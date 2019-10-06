package main

import (
	"github.com/wonderivan/logger"
	"HuanKeServer/src/conf"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"HuanKeServer/src/util"
	"HuanKeServer/src/controller"
	"HuanKeServer/src/model"
)

func main(){
	if !InitServer() {
		logger.Error("init server error!")
		return
	}
	logger.Debug("huanke server run start")

	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", configration.Global.Mysql.AdminDsn, 30)
	orm.RegisterModel(new(model.TsUser))

	beego.BConfig.RunMode = "prod"
	beego.BConfig.RouterCaseSensitive = true
	beego.BConfig.EnableErrorsShow = false
	beego.BConfig.EnableErrorsRender = false
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true

	beego.Router("/huanke/index", &controller.IndexController{})
	beego.Router("/huanke/register", &controller.RegisterController{})

	beego.Run(configration.Global.Server.Bindaddr)
}

func InitServer() bool {
	var currentPath = util.GetCurrentDirectory()
	var parentPath = util.GetParentDirectory(currentPath)
	var logConfigPath = parentPath + "/conf/log.json"
	var configPath = parentPath + "/conf/config.xml"

	//init log setting
	err := logger.SetLogger(logConfigPath)
	if err!=nil {
		log.Fatal(err)
		return false
	}

	err = configration.InitData(configPath)
	if err!=nil {
		logger.Error(err)
		return false
	}
	return true
}