package controllers

import (
	"fmt"
	"github.com/gaodongxu1/Beego/application/sample/models"
	"github.com/astaxie/beego"
	"strconv"
	"encoding/json"
	"log"
)

type Testcontrollers struct {
	beego.Controller
}

type user struct {
    Id    int         `form:"id"`
    Name  interface{} `form:"username"`
    Age   int         `form:"age"`
    Email string   
}

type student struct {
	Id int           `form:"id"`
	Name interface{} `form:"name"`
	age int          `form:"age"`
}

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

func (t *Testcontrollers) InsertMulti() {
	err := models.Test.InsertMulti()
	//fmt.Println(id, err)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"content": err}
		goto finish
	}
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

func(t *Testcontrollers) Updata() {
	 err := models.Test.Updata()
	if err!=nil {
		t.Data["json"] = map[string]interface{}{"content":err}
		goto finish
	}
	finish:
	 t.ServeJSON()
}

func(t *Testcontrollers) Delete() {
	err :=models.Test.Delete()
	if err !=nil {
		t.Data["json"] =map[string] interface{}{"content":err}
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
 // 获取int类型的数据
func (t *Testcontrollers) Getid() {
	fmt.Println("Getid........Test")
    id := t.Input().Get("id")
	intid, _ := strconv.Atoi(id)
	fmt.Printf("get the data  %d\n",intid)
}

// 把表单里的内容赋值到struct里
func (t *Testcontrollers) User() {
	u := user{}
    if err := t.ParseForm(&u); err != nil {
		//handle error
		t.Data["json"] = map[string]interface{}{"content":err}
		goto finish
	}
	t.Data["json"] = map[string]interface{}{"content":u}

	finish:
	t.ServeJSON()

	fmt.Printf("get the data  %v\n",u)
}
func(t *Testcontrollers) Student() {
		 s :=student{}
		 if err :=t.ParseForm(&s);err != nil {
			fmt.Printf("get the data  %v\n",err)
			  t.Data["json"] = map[string]interface{}{"content":err}
			 goto finish
		 }
		 t.Data["json"]=map[string]interface{}{"content":s}
		 fmt.Printf("get the data  %v\n",s)
		 finish:
		 t.ServeJSON()
}

// 在 beego 中获取 Request Body 里的 JSON 或 XML 的数据
func (t *Testcontrollers) RequestBody() {
	var ob user
   err := json.Unmarshal(t.Ctx.Input.RequestBody, &ob)
   if err != nil {
	  fmt.Printf("get the data  %v\n",err)
	  t.Data["json"] = map[string]interface{}{"content":err}
       goto finish
   }
    fmt.Println("us", ob)
	t.Data["json"] =  map[string]interface{}{"content":ob}
	finish:
    t.ServeJSON()
}
// 上传文件
func (t *Testcontrollers) FileUpload() {
	fmt.Println("up")
    f, h, err := t.GetFile("user")
    if err != nil {
        log.Fatal("getfile err ", err)
    }
    defer f.Close()
    t.SaveToFile("user", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
}
// 数据绑定（?id=123&isok=true&ft=1.2&ol[0]=1&ol[1]=2&ul[]=str&ul[]=array&user.Name=astaxie）
func (t *Testcontrollers) BindData() {
	var id int
	t.Ctx.Input.Bind(&id, "id")  //id ==123
	
	var isok bool
	t.Ctx.Input.Bind(&isok, "isok")  //isok ==true
	
	var ft float64
	t.Ctx.Input.Bind(&ft, "ft")  //ft ==1.2
	
	ol := make([]int, 0, 2)
	t.Ctx.Input.Bind(&ol, "ol")  //ol ==[1 2]
	
	ul := make([]string, 0, 2)
	t.Ctx.Input.Bind(&ul, "ul")  //ul ==[str array]
	
	var user struct{Name  string}
	t.Ctx.Input.Bind(&user, "user")  //user =={Name:"astaxie"}
	fmt.Println("aaaa \n",id,isok,ft,ol,ul,user)
} 
// session使用
func(t *Testcontrollers) Session() {
	t.SetSession("user","gfdsff")
	fmt.Println(t.GetSession("user").(string))
}