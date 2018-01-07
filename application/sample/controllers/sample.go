package controllers

import (
	"fmt"
	"github.com/gaodongxu1/Beego/application/sample/models"
	"github.com/astaxie/beego"
	"strconv"
)

type Testcontrollers struct {
	beego.Controller
}

type user struct {
    Id    int         `form:"-"`
    Name  interface{} `form:"username"`
    Age   int         `form:"age"`
    Email string
}
// func (o *Sample) HelloWorld() {
// 	o.Data["json"] = map[string]string{"content": models.HelloWorld("Gao")}
// 	o.ServeJSON()
// }
// func (o *Sample) Nba() {
// 	o.Data["json"] = map[string] string{"content": models.Nba()}
// 	o.ServeJSON()
// }
// func (o *Sample) Test() {
// 	o.Data["json"] = map[string] string{"content": "hello Beego 正则路由"}
// 	o.ServeJSON()
// }
 func (t *Testcontrollers) Hello() {
 	t.Data["json"] = map[string] string{"content": "hello Beego "}
 	t.ServeJSON()
 }
func (t *Testcontrollers) Insert() {
	id, err := models.Test.Insert()
	fmt.Println(id, err)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"content": err}
		goto finish
	}
	t.Data["json"] = map[string]interface{}{"id": id}
finish:
	t.ServeJSON()
}

func (t *Testcontrollers) Read() {
	err := models.Test.Read()
	if err != nil {
		t.Data["json"] = map[string]interface{}{"content": err}
		goto finish
	}
finish:
	t.ServeJSON()
}

func (t *Testcontrollers) Params() {
	jsoninfo := t.GetString("a")
    if jsoninfo == "" {
        t.Ctx.WriteString("a is empty")
        return
    } else {
		t.Ctx.WriteString(jsoninfo)
        return
	}
}

func (t *Testcontrollers) Getid() {
	fmt.Println("Getid........Test")
    id := t.Input().Get("id")
	intid, _ := strconv.Atoi(id)
	fmt.Printf("get the data  %d\n",intid)
}

func (t *Testcontrollers) User() {
	u := user{}
    if err := t.ParseForm(&u); err != nil {
		//handle error
	}
	fmt.Printf("get the data  %d\n",t.ParseForm(&u))
}