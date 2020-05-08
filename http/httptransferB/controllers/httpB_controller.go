package controllers

import (
	"fmt"
)

type HttpB_controller struct {
	ApiController
}

// @Description wenjianliu 做的时候记得把登录注释了*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /httpController [post]
func (this *HttpB_controller) HttpController() {
	fmt.Println("----------进入---------")
	var form Zhi
	if !this.ParseForm(&form) {
		fmt.Println("进入参数退出")
		return
	}
	fmt.Println(form)
	this.Data["json"] = form
	this.ServeJSON()
	this.err = nil
	fmt.Println("----------结束---------")
}
