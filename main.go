package main

import (
	"encoding/gob"
	"encoding/json"
	_ "pet/docs"
	"pet/models"
	_ "pet/routers"
	"web"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

var dbAddress string = "root:38143195@tcp(192.168.33.11:3306)/pet?charset=utf8"

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", dbAddress)
	orm.Debug = true
	beego.SetLevel(beego.LevelInformational)
	beego.SessionOn = true
	beego.SessionProvider = "file"
	beego.SessionSavePath = "./tmp"
	beego.SessionGCMaxLifetime = 60 * 60 * 60 * 24

	gob.Register(models.Users{})

}
func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.SetStaticPath("/doc", "static/swagger")
	}
	beego.EnableDocs = true

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Run()
}

var FilterUser = func(ctx *context.Context) {
	user := ctx.Input.Session("user")
	if user == nil && ctx.Request.URL.Path != "/v1/users/login" && ctx.Request.URL.Path[:4] != "/doc" {
		outPut := helper.Reponse(1, nil, "请先登录")
		b, _ := json.Marshal(outPut)
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.WriteString(string(b))
	}
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
}
