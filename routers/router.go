package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/login", &controllers.UserController{}, "get:Login")
	beego.Router("/user/dologin", &controllers.UserController{}, "post:DoLogin")
    beego.Router("/user/reg", &controllers.UserController{}, "get:Reg")
    beego.Router("/user/doreg", &controllers.UserController{}, "post:DoReg")
}
