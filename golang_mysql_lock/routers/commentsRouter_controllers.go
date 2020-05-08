package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["golang_mysql_lock/controllers:CeshisController"] = append(beego.GlobalControllerRouter["golang_mysql_lock/controllers:CeshisController"],
		beego.ControllerComments{
			Method:           "Wenjianliu",
			Router:           `/wenjianliu`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_mysql_lock/controllers:GolangMysqlLock"] = append(beego.GlobalControllerRouter["golang_mysql_lock/controllers:GolangMysqlLock"],
		beego.ControllerComments{
			Method:           "LockOptimi",
			Router:           `/lock_optimi`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_mysql_lock/controllers:GolangMysqlLock"] = append(beego.GlobalControllerRouter["golang_mysql_lock/controllers:GolangMysqlLock"],
		beego.ControllerComments{
			Method:           "LockPessimi",
			Router:           `/lock_pessimi`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_mysql_lock/controllers:GolangMysqlLock"] = append(beego.GlobalControllerRouter["golang_mysql_lock/controllers:GolangMysqlLock"],
		beego.ControllerComments{
			Method:           "LockPessimiFor",
			Router:           `/lock_pessimi_for`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
