package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
	"time"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	Addtime   int64
	Logintime int64
}

func init() {
	orm.RegisterModel(new(User))
	var mysql_host string = beego.AppConfig.String("mysql_host")
	var mysql_port string = beego.AppConfig.String("mysql_port")
	var mysql_user string = beego.AppConfig.String("mysql_user")
	var mysql_pass string = beego.AppConfig.String("mysql_pass")
	var mysql_db string = beego.AppConfig.String("mysql_db")
	//var link string = mysql_user + ":" + mysql_pass + "@tcp(" + mysql_host + ":" + mysql_port + ")/" + mysql_db + "?charset=utf8"
	var link string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysql_user, mysql_pass, mysql_host, mysql_port, mysql_db)
	fmt.Println(link)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", link)
	orm.RunSyncdb("default", false, true)
}

func AddUser(username string, password string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := new(User)
	user.Username = username
	user.Password = MD5(password)
	user.Addtime = time.Now().Unix()
	user.Logintime = time.Now().Unix()
	//user := User{Username:username, Password:MD5(password), Addtime:time.Now().Unix(), Logintime:time.Now().Unix()}
	id, err := o.Insert(user)
	return id, err
}

func FindUser(username string, password string) User {
	o := orm.NewOrm()
	o.Using("default")
	var user User
	o.QueryTable("user").Filter("username", username).Filter("password", MD5(password)).One(&user)
	return user
}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
