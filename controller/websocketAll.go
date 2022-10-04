package controller

import (
	"encoding/json"
	"gin_websocket_test/MyJwt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader指定升级HTTP连接为WebSocket连接。
var upgraderAll = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var wsMapAll = make(map[string]*websocket.Conn)

type MessageAllStruct struct {
	Message  string `json:"message"` //内容
	Username string `json:"username"`
	Time     string `json:"time"`
}

// 发送接收消息
func WebsocketAll(c *gin.Context) {

	// 1.获取句柄
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

	// 2.解析参数
	token := c.Query("token")
	if token == "" {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "token不能为空",
		})
		c.Abort()
		return
	}
	claims, _ := MyJwt.ParseToken(token)

	// 3.存入map集合
	wsMapAll[claims.Identity] = conn

	// 4.监听接收消息
	for {
		//拿到发送过来的数据
		messaageInfo := new(MessageStruct)
		err := conn.ReadJSON(messaageInfo) //readjson格式化到结构体 ReadMessage是原始的 单测用的message接收的可以去看 发送可以用message
		if err != nil {
			log.Println("拿到发送过来的数据异常", err.Error())
			return
		}

		// 5.发送消息
		for _, conn := range wsMapAll {
			messaageInfo.Username = claims.Username
			messaageInfo.Time = time.Now().Format("2006-01-02 15:04:05")
			resJson, _ := json.Marshal(messaageInfo)
			err := conn.WriteMessage(websocket.TextMessage, resJson)
			if err != nil {
				log.Println("发送数据失败", err.Error())
				return
			}
		}
	}

}
