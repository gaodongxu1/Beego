
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)
type Testmodels struct{}

var Test *Testmodels

type Carts struct {
	Id      int
	Title   string
	Content string
}


// 注册模型
func init() {
	orm.RegisterModel(new(Carts))
}

func (insert *Testmodels) Insert() (int64, error) {
	o := orm.NewOrm() // 创建一个 Ormer
	// NewOrm 的同时会执行 orm.BootStrap (整个 app 只执行一次)，用以验证模型之间的定义并缓存。
	o.Using("test") // 默认使用 default，你可以指定为其他数据库

	carts := new(Carts)
	carts.Title = "my first"
	carts.Content = "qwertyui"
	fmt.Println(carts)
	id, err := o.Insert(carts)
	if err != nil {
		fmt.Println("错误：",err)
		return 0, err
	}
	fmt.Println(1234)
	return id, err
}

//插入多个
func(InsertMulti *Testmodels) InsertMulti() error{
	o :=orm.NewOrm()
	o.Using("test")
	carts := []Carts{
		{Title: "slene", Content:"qwer"},
		{Title: "astaxie",Content:"asdf"},
		{Title: "unknown",Content:"zxcv"},
	}
	_, err := o.InsertMulti(3, carts)
	//第一个参数 bulk 为并列插入的数量，第二个为对象的slice返回值为成功插入的数量
	return err
}


func (read *Testmodels) Read() error {
	o := orm.NewOrm()
	carts := Carts{Id: 124}
	
	err := o.Read(&carts)
	
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(carts.Id, carts.Title)
	}
	return err
}

func(updata *Testmodels) Updata() error {
	o := orm.NewOrm()
	o.Using("test")
	var carts Carts = Carts{
		Id: 123,
	}
	if o.Read(&carts) == nil {
		fmt.Printf("qqqqqq  %d\n", &carts)
		carts.Title = "MyName"
		if _, err := o.Update(&carts); err != nil {
			return err
		}
	}
	return nil
}
func(delete *Testmodels) Delete() error {
	o := orm.NewOrm()
	if _, err := o.Delete(&Carts{Id: 123}); err == nil {
		return err
	}
	return nil
}
