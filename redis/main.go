package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"redis/controllers"
	_ "redis/model"
	_ "redis/routers"
)

func main() {

	beego.SetStaticPath("/swagger", "swagger")
	beego.SetStaticPath("/file", "file")
	beego.ErrorController(&controllers.GlobalErrorController{})

	beego.Run()
}
