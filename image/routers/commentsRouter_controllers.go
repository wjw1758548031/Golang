package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["image/controllers:CeshisController"] = append(beego.GlobalControllerRouter["image/controllers:CeshisController"],
		beego.ControllerComments{
			Method:           "Wenjianliu",
			Router:           `/wenjianliu`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["image/controllers:ImageController"] = append(beego.GlobalControllerRouter["image/controllers:ImageController"],
		beego.ControllerComments{
			Method:           "TestImageA",
			Router:           `/testImageA`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
