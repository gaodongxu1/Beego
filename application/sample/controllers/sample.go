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
func (o *Sample) Nba() {
	o.Data["json"] = map[string] string{"content": models.Nba()}
	o.ServeJSON()
}
func (o *Sample) Test() {
	o.Data["json"] = map[string] string{"content": "hello Beego"}
	o.ServeJSON()
}