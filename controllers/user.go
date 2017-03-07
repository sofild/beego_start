package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"hello/models"
	"strconv"
	"time"
)

type UserController struct {
	beego.Controller
}

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

var globalSession *session.Manager

func init() {
	var sessionConfig = new(session.ManagerConfig)
	sessionConfig.CookieName = "gosessionid"
	sessionConfig.Gclifetime = 3600

	globalSession, _ = session.NewManager("memery", sessionConfig)
	//globalSession,_ = session.NewManager("memery", `{"cookieName":"gosessionid","gclifetime":3600}`)
	//go globalSession.GC()
}

func (c *UserController) Login() {
	c.TplName = "login.tpl"
}

func (c *UserController) DoLogin() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		beego.Info(err)
	} else {
		uinfo := models.FindUser(u.Username, u.Password)
		fmt.Println(uinfo)
		if uinfo.Id > 0 {
			uinfo.Logintime = time.Now().Unix()
			models.UpUser(uinfo)
            /*
			sess, _ := globalSession.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
			sess.Set("uid", uinfo.Id)
			sess.Set("username", uinfo.Username)
            */

            session := c.StartSession()
            session.Set("uid", uinfo.Id)

			c.Ctx.WriteString("UID:" + strconv.FormatInt(uinfo.Id, 10))
		} else {
			c.Ctx.WriteString("Login Failed.")
		}
	}
}

func (c *UserController) Reg() {
	var username string = "admin"
	var password string = "admin"
	uid, err := models.AddUser(username, password)
	uidStr := strconv.FormatInt(uid, 10)
	c.Ctx.WriteString(uidStr)
	fmt.Println(err)
}

func (c *UserController) DoReg() {
	fmt.Println("Hello")
}
