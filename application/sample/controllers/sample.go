package controllers

import (
	"github.com/gaodongxu1/Beego/application/sample/models"
	"github.com/astaxie/beego"
)

type Sample struct {
	beego.Controller
}


func (o *Sample) HelloWorld() {
	o.Data["json"] = map[string]string{"content": models.HelloWorld("Gao")}
	o.ServeJSON()
}
