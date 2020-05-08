package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["scs/controllers:CeshisController"] = append(beego.GlobalControllerRouter["scs/controllers:CeshisController"],
		beego.ControllerComments{
			Method:           "Wenjianliuz",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:CeshisController"] = append(beego.GlobalControllerRouter["scs/controllers:CeshisController"],
		beego.ControllerComments{
			Method:           "Wenjianliu",
			Router:           `/wenjianliu`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:WenJianLiu"] = append(beego.GlobalControllerRouter["scs/controllers:WenJianLiu"],
		beego.ControllerComments{
			Method:           "Filepath",
			Router:           `/filepath`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:WenJianLiu"] = append(beego.GlobalControllerRouter["scs/controllers:WenJianLiu"],
		beego.ControllerComments{
			Method:           "WenjianliuCreate",
			Router:           `/wenjianliuCreate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:WenJianLiu"] = append(beego.GlobalControllerRouter["scs/controllers:WenJianLiu"],
		beego.ControllerComments{
			Method:           "WenjianliuCreatexce",
			Router:           `/wenjianliuCreatexce`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:WenJianLiu"] = append(beego.GlobalControllerRouter["scs/controllers:WenJianLiu"],
		beego.ControllerComments{
			Method:           "WenjianliuExcelizeAddIconexce",
			Router:           `/wenjianliuExcelizeAddIconexce`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:WenJianLiu"] = append(beego.GlobalControllerRouter["scs/controllers:WenJianLiu"],
		beego.ControllerComments{
			Method:           "WenjianliuExcelizeCreatexce",
			Router:           `/wenjianliuExcelizeCreatexce`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:WenJianLiu"] = append(beego.GlobalControllerRouter["scs/controllers:WenJianLiu"],
		beego.ControllerComments{
			Method:           "WenjianliuExcelizeInsertImageexce",
			Router:           `/wenjianliuExcelizeInsertImageexce`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:WenJianLiu"] = append(beego.GlobalControllerRouter["scs/controllers:WenJianLiu"],
		beego.ControllerComments{
			Method:           "WenjianliuExcelizeReadexce",
			Router:           `/wenjianliuExcelizeReadexce`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:WenJianLiu"] = append(beego.GlobalControllerRouter["scs/controllers:WenJianLiu"],
		beego.ControllerComments{
			Method:           "WenjianliuUpdateexce",
			Router:           `/wenjianliuUpdateexce`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:WenJianLiu"] = append(beego.GlobalControllerRouter["scs/controllers:WenJianLiu"],
		beego.ControllerComments{
			Method:           "WenjianliuWrite",
			Router:           `/wenjianliuWrite`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:WenJianLiu"] = append(beego.GlobalControllerRouter["scs/controllers:WenJianLiu"],
		beego.ControllerComments{
			Method:           "WenjianliuWriteCreateA",
			Router:           `/wenjianliuWriteCreateA`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:WenJianLiu"] = append(beego.GlobalControllerRouter["scs/controllers:WenJianLiu"],
		beego.ControllerComments{
			Method:           "Wenjianliuexce",
			Router:           `/wenjianliuexce`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:WenJianLiu"] = append(beego.GlobalControllerRouter["scs/controllers:WenJianLiu"],
		beego.ControllerComments{
			Method:           "WenjianliutxtOpen",
			Router:           `/wenjianliutxtOpen`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["scs/controllers:WenJianLiu"] = append(beego.GlobalControllerRouter["scs/controllers:WenJianLiu"],
		beego.ControllerComments{
			Method:           "Wenjianliutxtt",
			Router:           `/wenjianliutxtt`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
