package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["httptransferB/controllers:CeshisController"] = append(beego.GlobalControllerRouter["httptransferB/controllers:CeshisController"],
		beego.ControllerComments{
			Method:           "Wenjianliu",
			Router:           `/wenjianliu`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["httptransferB/controllers:HttpB_controller"] = append(beego.GlobalControllerRouter["httptransferB/controllers:HttpB_controller"],
		beego.ControllerComments{
			Method:           "HttpController",
			Router:           `/httpController`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
