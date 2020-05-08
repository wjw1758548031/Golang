package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["redis/controllers:CeshisController"] = append(beego.GlobalControllerRouter["redis/controllers:CeshisController"],
		beego.ControllerComments{
			Method:           "Wenjianliu",
			Router:           `/wenjianliu`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["redis/controllers:Redis"] = append(beego.GlobalControllerRouter["redis/controllers:Redis"],
		beego.ControllerComments{
			Method:           "GetRedis",
			Router:           `/getRedis`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["redis/controllers:Redis"] = append(beego.GlobalControllerRouter["redis/controllers:Redis"],
		beego.ControllerComments{
			Method:           "SetRedis",
			Router:           `/setRedis`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
