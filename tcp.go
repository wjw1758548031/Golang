package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

//int64 序列化最大容量为8
var (
	headLen  = int64(8)
	bodyByte = make([]byte, 0)
)

//tcp双项通道
func main() {
	go ServerBase()
	time.Sleep(1 * time.Second) //网速卡的话自己调整，可能在没创建完就已经连接
	go ClientBase()
	time.Sleep(1 * time.Hour)
}

func ServerBase() {
	fmt.Println("Starting the server...")
	//create listener
	listener, err := net.Listen("tcp", "127.0.0.1:50000")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	// listen and accept connections from clients:
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}
		//create a goroutine for each request.
		go doServerStuff(conn)
	}

}

type TcpDate struct {
	Header int64
	Body   string
}

func NewTcpDate(Body string) TcpDate {
	return TcpDate{Header: int64(len(Body)), Body: Body}
}

func (t TcpDate) EncodeHeader() []byte {
	b_buf := new(bytes.Buffer)
	if err := binary.Write(b_buf, binary.BigEndian, &t.Header); err != nil {
		panic(err)
	}
	return b_buf.Bytes()
}

func (t TcpDate) Encode() []byte {
	return []byte(string(t.EncodeHeader()) + t.Body)
}

func (t *TcpDate) Decode(data []byte, endLen int64) (bool, []byte) {
	if endLen > int64(len(data)-1) {
		return false, data
	}
	head := data[:endLen]
	b_buf := new(bytes.Buffer)
	b_buf.Write(head)
	if err := binary.Read(b_buf, binary.BigEndian, &t.Header); err != nil {
		panic(err)
	}
	body := data[endLen : endLen+t.Header]
	t.Body = string(body)
	fmt.Println(t.Body)
	return true, data[endLen+t.Header:]
}

func doServerStuff(conn net.Conn) {
	fmt.Println("new connection:", conn.LocalAddr())
	for {
		data := TcpDate{}
		//自己定义长度，如果数据过多。。。。
		buf := make([]byte, 1024)
		bufLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		bodyByte = []byte(fmt.Sprintf("%v%v", string(bodyByte), string(buf[:bufLen])))
		for {
			var is bool
			is, bodyByte = data.Decode(bodyByte, headLen)
			if !is {
				break
			}
		}
		/*len,_ := strconv.Atoi(strings.TrimPrefix(string(head),"ceshi:"))
		fmt.Println("data:",string(buf[100:len]))*/
		conn.Write([]byte("已收到"))

	}
}

func ClientBase() {
	//open connection:
	conn, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		fmt.Println("Error dial:", err.Error())
		return
	}
	//send info to server until Quit
	for i := 0; i < 100; i++ {
		go func() {
			item := struct {
				Name1 string
				Name2 string
				Name3 string
				Name4 string
				Name5 string
			}{}

			body, _ := json.Marshal(item)
			data := NewTcpDate(string(body))
			dataByte := data.Encode()
			_, err := conn.Write(dataByte)
			if err != nil {
				fmt.Println("Error Write:", err.Error())
				return
			}

			//这里就是简单的接收
			buf := make([]byte, 1024)
			length, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
				return
			}
			fmt.Println("Receive data from server:", string(buf[:length]))
		}()
	}
}
