// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/gaodongxu1/Beego/application/sample/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.Sample{}, "get:HelloWorld")
	// beego.Router("/nba", &controllers.Sample{}, "*:Nba")  //自定义方法
	// beego.Router("/api/123:id=123", &controllers.Sample{}, "get:Test") //正则路由 
	beego.Router("/hello",&controllers.Testcontrollers{}, "get:Hello")
	beego.Router("/insert", &controllers.Testcontrollers{}, "get:Insert")
	beego.Router("/insertmulti", &controllers.Testcontrollers{}, "get:InsertMulti")
	beego.Router("/read", &controllers.Testcontrollers{}, "get:Read")
	beego.Router("/params", &controllers.Testcontrollers{}, "post:Params")
	beego.Router("/getid", &controllers.Testcontrollers{}, "post:Getid")
	beego.Router("/user", &controllers.Testcontrollers{}, "get:User")
	beego.Router("/student", &controllers.Testcontrollers{}, "get:Student")
	beego.Router("/requestbody", &controllers.Testcontrollers{}, "post:RequestBody")
	beego.Router("/upload", &controllers.Testcontrollers{}, "post:FileUpload")
	beego.Router("/bind", &controllers.Testcontrollers{}, "get:BindData")
	beego.Router("/updata", &controllers.Testcontrollers{}, "get:Updata")
	beego.Router("/delete", &controllers.Testcontrollers{}, "get:Delete")
	beego.Router("/session", &controllers.Testcontrollers{}, "get:Session")
}
