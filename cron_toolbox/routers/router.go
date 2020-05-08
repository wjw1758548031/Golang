// @Title 物供系统接口文档
// @Description 物供系统所有的接口都会在该文档中进行说明
package routers

import (
	"cron_toolbox/controllers"
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

func init() {

	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")
	beego.Router("/sendMessage", &controllers.WebSocketController{}, "post:SendMessage")
	ns1 := beego.NewNamespace("/1",
		beego.NSNamespace("/ceshis",
			beego.NSInclude(
				&controllers.CeshisController{},
			)),
		beego.NSNamespace("/toolbox",
			beego.NSInclude(
				&controllers.Toolbox{},
			)),
	)
	beego.AddNamespace(ns1)

	beego.Handler("/captcha/*.png", captcha.Server(120, 40)) //注册验证码服务，验证码图片的宽高为240 x 80
}
