package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["cron_toolbox/controllers:CeshisController"] = append(beego.GlobalControllerRouter["cron_toolbox/controllers:CeshisController"],
		beego.ControllerComments{
			Method:           "Wenjianliu",
			Router:           `/wenjianliu`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["cron_toolbox/controllers:Toolbox"] = append(beego.GlobalControllerRouter["cron_toolbox/controllers:Toolbox"],
		beego.ControllerComments{
			Method:           "Toolbox",
			Router:           `/toolbox`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
