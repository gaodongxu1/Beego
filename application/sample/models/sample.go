// package models

// import (
// 	"fmt"
// )

// func HelloWorld(name string) string {
// 	return fmt.Sprintf("Hello %s!", name)
// }
// func Nba() string {
// 	return fmt.Sprintf("NBA直播吧")
// }
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// type MessageServiceProvider struct{}

// var MessageService *MessageServiceProvider

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
	carts.Title = "my first bolg"
	carts.Content = "默认使用"
	id, err := o.Insert(carts)
	if err != nil {
		return 0, err
	}

	return id, err
}
func (read *Testmodels) Read() error {
	o := orm.NewOrm()
	carts := Carts{Id: 1000}
	
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
