package controllers

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type WebSocket struct {
	ApiController
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{}

type Message struct {
	Message string `json:"message"`
}

func init() {
	/*fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	//建立连接
	http.HandleFunc("/ws", handleConnections)
	*/ //使用beego.run() 则不需要创建连接
	//启动线程循环
	go handleMessages()

	//log.Println("http server started on :8501")
	//开启8000端口
	/*err := http.ListenAndServe(":8501", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}*/
}

//注册成为 websocket 创建连接
func (this *WebSocket) HandleConnections() {
	w := this.Ctx.ResponseWriter
	r := this.Ctx.Request
	log.Println("进入handleConnections")
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	clients[ws] = true
	log.Println("clients:", clients)

	//不断的从页面上获取数据 然后广播发送出去
	for {
		//将从页面上接收数据改为不接收 直接发送
		//var msg Message
		//err := ws.ReadJSON(&msg)
		//if err != nil {
		//  log.Printf("error: %v", err)
		//  delete(clients, ws)
		//  break
		//}

		//目前存在问题 定时效果不好 需要在业务代码替换时改为beego toolbox中的定时器 需要定时器触发
		time.Sleep(time.Second * 3)
		msg := Message{Message: "这是向页面发送的数据 " + time.Now().Format("2006-01-02 15:04:05")}
		broadcast <- msg
	}

}

// @Description 发送推送*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /toSendMessage [post]
func (this *WebSocket) ToSendMessage() {
	log.Println("进入toSendMessage--11--")
	time.Sleep(time.Second * 3)
	msg := Message{Message: "这是向页面发送的数据 " + time.Now().Format("2006-01-02 15:04:05")}
	broadcast <- msg
}

//广播发送至页面
func handleMessages() {
	for {
		log.Println("clients:", clients)
		msg := <-broadcast
		for client := range clients {
			log.Println("推送:", msg)
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("client.WriteJSON error: %v", err)
				client.Close()
				delete(clients, client)

			}
		}
	}
}

//前端react代码

/*

import React from 'react';
import '../assets/css/Whole.css'

class Goods extends  React.Component {

    constructor(props){
        //用户父子组件传值
        super(props)

        this.state = {
            title : "标签"
        }

        alert("ss")
        var ws = new WebSocket('ws://' + 'localhost:8501' + '/ws');
        ws.onmessage = function(e) {
           alert(JSON.stringify(e.data))
        };
        alert("ss")
    }

    websocket() {
        alert("进入 websocket")
        var init = {
            body: JSON.stringify({}), // must match 'Content-Type' header
            cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
            headers: {
                'content-type': 'text/plain'
            },
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            mode: 'cors', // no-cors, cors, *same-origin
        }
        fetch(
            'http://localhost:8501/1/web_socket/toSendMessage',
             init
        )
            .then(res =>   alert("请求成功") /*res.json()*/ //)
/*
.then(data => {
alert("请求成功")
})
.catch(e => console.log('错误:', e))
}


render(){
return (<div>

<div onClick={this.websocket}>推送</div>

<div className="red" title={this.state.title+"1"}>商品</div>


<div  title={this.state.title+"2"}>商品2</div>

<div for="name">点击我选中文本框</div>

<input id="name" />


</div>)
}

}

export default Goods;

*/
