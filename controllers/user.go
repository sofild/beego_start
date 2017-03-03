package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
    "hello/models"
    "strconv"
)

type UserController struct {
	beego.Controller
}

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (c *UserController) Login() {
	c.TplName = "login.tpl"
}

func (c *UserController) DoLogin() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		beego.Info(err)
	} else {
	    uinfo := models.FindUser(u.username,u.password)
        if uinfo==nil {
            c.Ctx.WriteString("Login Failed.")
        }
        else{
            c.Ctx.WriteString("UID:"+uinfo.Id)
        }
    }
}

func (c *UserController) Reg(){
    var username string = "admin"
    var password string = "admin"
    uid,err := models.AddUser(username,password)
    uidStr := strconv.FormatInt(uid, 10)
    c.Ctx.WriteString(uidStr)
    fmt.Println(err)
}

func (c *UserController) DoReg(){
    fmt.Println("Hello")
}
