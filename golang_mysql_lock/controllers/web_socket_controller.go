package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"net/http"
)

// 定义websocket消息显示功能
type WebSocketController struct {
	beego.Controller
}

// 加入websocket聊天
func (this *WebSocketController) Join() {
	// 升级为websocket对象

	beego.Informational("客户端申请链接")
	uname := this.GetString("acs_token")
	logs.Info("***客户：" + uname + "加入websocket聊天室***")
	if ok := this.checkUser(uname); !ok {
		return
	}

	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "无效的websocket协议", 400)
		return
	} else if err != nil {
		http.Error(this.Ctx.ResponseWriter, "不能建立起websocket连接:", 400)
		return
	}

	// 循环读取消息
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			beego.Error("服务端读取信息错误:", err)
			return
		}
		// 正常读取到消息，则发布一个message消息
		//Publish <- ActionForm{"wjw", "203", LoginUserInfo{Name:"普通用户"}, WebSocketVo{Content:string(p)}, DependenceForm{}, nil}
	}
}

func (this *WebSocketController) SendMessage() {
	//ToSendMessage("111")
}

func (this *WebSocketController) checkUser(acsToken string) bool {
	/*	_, err := GetLoginUser(acsToken)
		if err != nil {
			errResponse := ErrResponseVo{ErrCode: "230", ErrMsg: err.Error()}
			this.Ctx.Output.JSON(errResponse, false, true)
			return false
		}*/
	return true
}
