package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
	"httptransferB/controllers"
	_ "httptransferB/model"
	_ "httptransferB/routers"
)

func main() {

	/*beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	        AllowOrigins: []string{"https://*"},
	        AllowMethods: []string{"PUT", "PATCH"},
	        AllowHeaders: []string{"Origin"},
	        ExposeHeaders: []string{"Content-Length"},
	        AllowCredentials: true,    }))

	    beego.Run()
	}
	*/

	/* beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	     AllowOrigins:     []string{"https://*"},
	     AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	     AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	    ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	     AllowCredentials: true,
	}))*/

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	beego.SetStaticPath("/swagger", "swagger")
	beego.SetStaticPath("/file", "file")
	beego.ErrorController(&controllers.GlobalErrorController{})
	beego.Run()
}
