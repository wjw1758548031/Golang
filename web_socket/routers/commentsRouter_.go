package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["mutex/controllers:CeshisController"] = append(beego.GlobalControllerRouter["mutex/controllers:CeshisController"],
		beego.ControllerComments{
			Method:           "Wenjianliu",
			Router:           `/wenjianliu`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mutex/controllers:Mutex"] = append(beego.GlobalControllerRouter["mutex/controllers:Mutex"],
		beego.ControllerComments{
			Method:           "Ceshi",
			Router:           `/ceshi`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mutex/controllers:Mutex"] = append(beego.GlobalControllerRouter["mutex/controllers:Mutex"],
		beego.ControllerComments{
			Method:           "Mutex",
			Router:           `/mutex`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mutex/controllers:Mutex"] = append(beego.GlobalControllerRouter["mutex/controllers:Mutex"],
		beego.ControllerComments{
			Method:           "Rwmutexx",
			Router:           `/rwmutexx`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
