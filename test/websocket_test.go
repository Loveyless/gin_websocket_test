package test

import (
	"flag"
	"log"
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
)

//
var addr = flag.String("addr", "localhost:8080", "http service address")

// Upgrader指定升级HTTP连接为WebSocket连接。
// 并发调用Upgrader的方法是安全的。
var upgrader = websocket.Upgrader{}

//一对多 定义一个ws
var wsList = make([]*websocket.Conn, 0)

func echo(w http.ResponseWriter, r *http.Request) {

	// Conn 类型表示一个 WebSocket 连接。服务器应用程序从 HTTP 请求处理程序调用 Upgrader.Upgrade 方法以获取 *Conn：
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()

	//如果有人连接 放到切片里
	wsList = append(wsList, conn)

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		//循环切片 一个一个发送 一对多 上面的逻辑先别管
		for _, conn := range wsList {
			conn.WriteMessage(mt, message)
		}
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func TestWebsocket(t *testing.T) {
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
