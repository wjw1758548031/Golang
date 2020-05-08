package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"httptransfer/server"
)

type CeshisController struct {
	ApiController
}

type Zhi struct {
	BoardId string `json:"board_id"  description:"角色ID"`
	Id      int    `json:"id"  description:"角色ID"`
}

/*
func (baseApi *ApiController) ParseForm(form interface{}) bool {
	err := json.Unmarshal(baseApi.Ctx.Input.RequestBody, form); if err != nil {
		utils.PrtErrError("JSON转换出错", err)
		baseApi.err = errorRequestJsonFormat
		return false
	}
	err = validForm(form); if err != nil {
		baseApi.err = err
		return false
	}
	return true
}*/

type AdviceForm struct {
	HandleType int `json:"handle_type" description:"1供应商通知 2客户通知" `
}

// @Description wenjianliu 做的时候记得把登录注释了*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliu [post]
func (this *CeshisController) Wenjianliu() {

	var adviceForm AdviceForm
	adviceForm.HandleType, _ = this.GetInt("handle_type")

	//response.setHeader("Access-Control-Allow-Origin", "*");
	fmt.Println("进入query")
	//input.Request.Form.Get(key)
	fmt.Println("this.Ctx.Input.RequestBody", this.Ctx.Input.RequestBody)
	/*var form Zhi
	if !this.ParseForm(&form) {
		fmt.Println("进入参数退出")
		return
	}*/

	o := orm.NewOrm()
	goods := []server.Goods{}
	_, err := o.Raw("select * from t_goods").QueryRows(&goods)
	if err != nil {
		fmt.Println("查询出错", err)
		return
	}
	fmt.Println("goods集合", goods)

	var rw = this.Ctx.ResponseWriter
	//允许访问所有域
	//rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:8501")
	//rw.Header().Set("Origin", "*")
	//rw.Header().Set("Authorization", "http://localhost:8501")
	//Access-Control-Allow-Headers
	// 设置Content-Type
	rw.Header().Set(`Content-Type`, `text/event-stream;charset=utf-8`)
	rw.Header().Set("Cache-Control", "no-cache")
	data, _ := json.Marshal(goods)
	var bs = bytes.NewBuffer(data)
	rw.Write(bs.Bytes())
	rw.Flush()

	//this.Data["json"] = goods
	/*this.Data["json"] = "123456"
	this.ServeJSON()*/
	//this.data =  "123456"
	this.Data["json"] = "123456"
	this.ServeJSON()
	this.err = nil
}
