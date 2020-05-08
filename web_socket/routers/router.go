// @Title 物供系统接口文档
// @Description 物供系统所有的接口都会在该文档中进行说明
package routers

import (
	"github.com/astaxie/beego"
	"web_socket/controllers"
)

func init() {

	beego.Router("/ws", &controllers.WebSocket{}, "get:HandleConnections")
	ns1 := beego.NewNamespace("/1",
		beego.NSNamespace("/web_socket",
			beego.NSInclude(
				&controllers.WebSocket{},
			)),
	)
	beego.AddNamespace(ns1)
}
