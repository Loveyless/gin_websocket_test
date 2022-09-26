package controller

import (
	"fmt"
	"gin_websocket_test/MyJwt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader指定升级HTTP连接为WebSocket连接。
// 并发调用Upgrader的方法是安全的。
var upgrader = websocket.Upgrader{}

//存储conn句柄的map key是用户id value是conn句柄
var wsMap = make(map[string]*websocket.Conn)

//接收发送消息的结构体
type MessageStruct struct {
	RoomIdentity string `json:"room_identity"` //接收房间
	Message      string `json:"message"`       //内容
}

// 发送接收消息
func WebsocketMessage(c *gin.Context) {

	//Upgrade将HTTP服务器连接升级到WebSocket协议。返回句柄
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "wensocket升级失败",
			"data":    err,
		})
		c.Abort()
		return
	}
	defer conn.Close()

	// token取出数据
	claims := c.MustGet("claims").(*MyJwt.MyCustomClaims)
	fmt.Println(claims.Identity)
	//存入map集合
	wsMap[claims.Identity] = conn

	//核心 for一直接收消息
	for {
		messaageInfo := new(MessageStruct)
		err := conn.ReadJSON(messaageInfo) //readjson格式化到结构体 ReadMessage是原始的 单测用的message接收的可以去看 发送可以用message
		if err != nil {
			log.Println("读取数据失败", err.Error())
			return
		}

		//发送消息 目前是给所有再线用户发 后面优化成给所有在线 并且在同一个房间的用户发
		for _, conn := range wsMap {
			//TextMessage表示文本数据消息。文本消息有效负载被解释为UTF-8编码的文本数据。
			//消息类型是一个整数，表示消息的类型，它的值可以是下面几个常量之一： TextMessage对应1、BinaryMessage、CloseMessage、PingMessage、PongMessage
			err := conn.WriteMessage(websocket.TextMessage, []byte(messaageInfo.Message))
			if err != nil {
				log.Println("发送数据失败", err.Error())
				return
			}
		}
	}

}
