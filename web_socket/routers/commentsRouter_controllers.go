package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["web_socket/controllers:WebSocket"] = append(beego.GlobalControllerRouter["web_socket/controllers:WebSocket"],
		beego.ControllerComments{
			Method:           "ToSendMessage",
			Router:           `/toSendMessage`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
