package main

import (
	_ "nohasslsapi/docs"
	_ "nohasslsapi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"

	_ "github.com/go-sql-driver/mysql"
)

// "github.com/davecgh/go-spew/spew"
func init() {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/nohassls")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//beego.InsertFilter("*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("*", beego.FinishRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
	}))

	beego.Run()
}

var FilterUser = func(ctx *context.Context) {

	ctx.Output.Header("Access-Control-Allow-Origin", "*")

	return

	//	// ensure its a get
	//	if ctx.Input.Method() != "GET" {
	//		ctx.ResponseWriter.WriteHeader(403)
	//		ctx.WriteString("Method Not Allowed. Only GET requests are supported")
	//		return
	//	}

	//	// simple grab app id
	//	var appid string

	//	if ctx.Input.Query("appid") == "" {
	//		// check header
	//		appid = ctx.Input.Header("appid")
	//	} else {
	//		// if not found check url string
	//		appid = ctx.Input.Query("appid")
	//	}

	//	//spew.Dump(appid)

	//	// if blank error
	//	if appid == "" {
	//		ctx.ResponseWriter.WriteHeader(400)
	//		ctx.WriteString("miss query param: appid")
	//		return
	//	}

	//	// check it matches our config.
	//	if appid == beego.AppConfig.String("appid") {
	//		ctx.Output.Header("version", beego.AppConfig.String("api_version"))
	//		ctx.Output.Header("Server", beego.AppConfig.String("server"))
	//		return
	//	}

	//	ctx.ResponseWriter.WriteHeader(403)
	//	var errmsglocal string
	//	errmsglocal = "appid_error_" + beego.AppConfig.String("region")
	//	ctx.WriteString(beego.AppConfig.String(errmsglocal))
}
