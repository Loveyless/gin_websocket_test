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
// 并发调用Upgrader的方法是安全的。
var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 在线用户的句柄 存储conn句柄的map
// key是用户Identity value是conn句柄
var wsMap = make(map[string]*websocket.Conn)

//接收发送消息的结构体
type MessageStruct struct {
	RoomIdentity string `json:"room_identity"` //接收房间
	Message      string `json:"message"`       //内容
	// Token        string `json:"token"`
	Username string `json:"username"`
	Time     string `json:"time"`
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

	// token取出数据 js里的websocket不能随便带请求头 只能自己解析
	// claims := c.MustGet("claims").(*MyJwt.MyCustomClaims)

	//自己从请求头解析 发现可以解析 但是
	token := c.GetHeader("Sec-WebSocket-Protocol")
	if token == "" {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "token不能为空",
		})
		c.Abort()
		return
	}
	claims, _ := MyJwt.ParseToken(token)

	//自己从参数解析
	// token := c.Query("tt")
	// if token == "" {
	// 	c.JSON(200, gin.H{
	// 		"status":  400,
	// 		"message": "token不能为空",
	// 	})
	// 	c.Abort()
	// 	return
	// }
	// claims, _ := MyJwt.ParseToken(token)

	// 请求头和参数都解析不出来 只能从发过来的地方解析了 所以下方的存入map集合 也做不了
	// 只能在拿到数据后存入map集合

	//存入map集合
	wsMap[claims.Identity] = conn

	//核心 for一直接收消息
	for {
		//拿到发送过来的数据
		messaageInfo := new(MessageStruct)
		err := conn.ReadJSON(messaageInfo) //readjson格式化到结构体 ReadMessage是原始的 单测用的message接收的可以去看 发送可以用message
		if err != nil {
			log.Println("拿到发送过来的数据异常", err.Error())
			return
		}

		// 1.判断用户是否属于消息体的房间
		// _, err = service.GetUserRoomByUserIdentityRoomIdentity(claims.Identity, messaageInfo.RoomIdentity)
		// if err != nil {
		// 	log.Printf("user_identity:%v,room_identity:%v error:%v", claims.Identity, messaageInfo.RoomIdentity, err.Error())
		// 	return
		// }
		// // 2.获取在特定房间的在线用户
		// // 	查询房间的所有用户
		// userRoomList, err := service.GetUserRoomByRoomIdentity(messaageInfo.RoomIdentity)
		// if err != nil {
		// 	log.Printf("room_identity:%v error:%v", messaageInfo.RoomIdentity, err.Error())
		// 	return
		// }

		//	遍历出特定房间的在线用户
		// for _, userRoom := range userRoomList {
		// 	// 如果房间中所有用户有在线的用户那么就发送消息 注意这里巧妙的写法!!
		// 	if conn, ok := wsMap[userRoom.UserIdentity]; ok {
		// 		err := conn.WriteMessage(websocket.TextMessage, []byte(messaageInfo.Message))
		// 		if err != nil {
		// 			log.Println("发送数据失败", err.Error())
		// 			return
		// 		}
		// 	}
		// }
		// 3.保存消息
		// messageBasic := &service.MessageBasic{
		// 	UserIdentity: claims.Identity,
		// 	RoomIdentity: messaageInfo.RoomIdentity,
		// 	Data:         messaageInfo.Message,
		// 	CratedAt:     time.Now().Unix(),
		// 	UpdatedAt:    time.Now().Unix(),
		// }
		// err = service.InsertOneMessageBasic(messageBasic)
		// if err != nil {
		// 	log.Printf("保存消息错误%v", err.Error())
		// 	return
		// }

		// //发送消息 目前是给所有再线用户发 后面优化成给所有在线 并且在同一个房间的用户发
		for _, conn := range wsMap {
			//TextMessage表示文本数据消息。文本消息有效负载被解释为UTF-8编码的文本数据。
			//消息类型是一个整数，表示消息的类型，它的值可以是下面几个常量之一： TextMessage对应1、BinaryMessage、CloseMessage、PingMessage、PongMessage
			// err := conn.WriteMessage(websocket.TextMessage, []byte(messaageInfo.Message))
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
