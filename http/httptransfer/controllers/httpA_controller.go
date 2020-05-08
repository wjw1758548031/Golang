package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpA_controller struct {
	ApiController
}

// @Description wenjianliu 做的时候记得把登录注释了*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /httpController [post]
func (this *HttpA_controller) HttpController() {
	fmt.Println("----------进入1---------")

	body, _ := json.Marshal(Zhi{BoardId: "9999", Id: 999})
	res, err := http.Post("http://localhost:8504/1/httpcontroller/httpController", "application/json;charset=utf-8", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	fmt.Println(string(content))

	fmt.Println("----------结束---------")
}
