package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["httptransfer/controllers:CeshisController"] = append(beego.GlobalControllerRouter["httptransfer/controllers:CeshisController"],
		beego.ControllerComments{
			Method:           "Wenjianliu",
			Router:           `/wenjianliu`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["httptransfer/controllers:HttpA_controller"] = append(beego.GlobalControllerRouter["httptransfer/controllers:HttpA_controller"],
		beego.ControllerComments{
			Method:           "HttpController",
			Router:           `/httpController`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
