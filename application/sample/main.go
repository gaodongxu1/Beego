package main

import (
	_ "github.com/gaodongxu1/Beego/application/sample/routers"

	"github.com/astaxie/beego"
	"github.com/gaodongxu1/Beego/application/sample/orm"
)

func main() {
	orm.InitMysql()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
