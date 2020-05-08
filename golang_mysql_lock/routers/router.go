// @Title 物供系统接口文档
// @Description 物供系统所有的接口都会在该文档中进行说明
package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"golang_mysql_lock/controllers"
)

func init() {

	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")
	beego.Router("/sendMessage", &controllers.WebSocketController{}, "post:SendMessage")
	ns1 := beego.NewNamespace("/1",
		beego.NSNamespace("/ceshis",
			beego.NSInclude(
				&controllers.CeshisController{},
			)),
		beego.NSNamespace("/golang_mysql_lock",
			beego.NSInclude(
				&controllers.GolangMysqlLock{},
			)),
	)
	beego.AddNamespace(ns1)

	beego.Handler("/captcha/*.png", captcha.Server(120, 40)) //注册验证码服务，验证码图片的宽高为240 x 80
}
