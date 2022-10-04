# websocket聊天室

## 前置准备
```
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/websocket
go get go.mongodb.org/mongo-driver/mongo
go get github.com/jordan-wright/email
```

**功能**

用户模块
密码登录
邮箱注册
用户详情
通讯模块（核心）
一对一通讯
多对多通讯
消息列表聊天记录列表


## 集合设计

**用户信息集合**

user

```Go
{
    "username":"账号",
    "password":"密码",
    "nicname":"用户名",
    "sex":1, //0-未知 1-男 2-女
    "email":"邮箱",
    "avatar":"头像",
    "created_at":1,
    "updated_at":1,
}
```



**消息集合**

message

```Go
{
    "user_identity":"用户唯一标识",
    "room_identity":"房间唯一标识",
    "data":"发送的数据",
    "created_at":1,
    "updated_at":1,
}
```



**房间集合**

room

```Go
{
    "identity":"房间唯一标识",
    "number":"房间号",
    "name":"房间名",
    "info":"简介",
    "user_identity":"房间创建者的唯一标识",
    "room_type":"房间类型",    //[1独聊] [2群聊]
    "created_at":1,
    "updated_at":1,
}
```



**房间和用户的关联**

user_room

```Go
{
    "user_identity": "用户的唯一标识",
    "room_identity": "房间的唯一标识",
    "message_identity": "消息的唯一标识",
    "created_at": 1,
    "updated_at": 1
}
```


## 目录结构

gin+mongodb

目录结构
```
|-config      基本配置
|-controller  基础逻辑
|-cors        跨域中间件
|-email       封装发送邮箱
|-MyJwt       token中间件
|-MyUtils     工具
|-router      路由 分组/拦截
|-service     操作数据的函数
    db.go     mongo redis数据库句柄
|-test        测试函数
|-validator   request校验
|-validator   将gin验证器错误翻译成中文
|-go_test.sql mongodb导出导入文件
|-go.mod
|-gosum
|-main.go
|-websocket项目.postman_collection.json postman导出导入文件
```





## 接口规则



普通接口规则在postman

websocket协议接口规则

```go
ws://localhost:8080/user/websocket/message
请求头
	token
接收发送消息的结构体
    type MessageStruct struct {
        RoomIdentity string `json:"room_identity"` //接收房间
        Message      string `json:"message"`       //内容
    }


新做了一个接口
ws://localhost:8080/websocket/all
请求参数
	token:xxx
接收发送消息的结构体
    type MessageStruct struct {
        Message      string `json:"message"`       //内容
    }
```



## 邮箱库的使用

### 发邮件

![](https://cdn.jsdelivr.net/gh/Loveyless/img-clouding/img/下载.png)

```Go
//NWNXHEUGDOYJJSYG
func TestSendEmail(t *testing.T) {
  e := email.NewEmail()                                                                                     //创建一个新的邮件
  e.From = "Loveyless <githubbyloveyless@163.com>"                                                          // 发件人
  e.To = []string{"59771463@qq.com"}                                                                        // 收件人 可以多个
  e.Subject = "验证码已发送，请查收"                                                                                  //邮件主题
  e.HTML = []byte("<div>Your verification code:<b>" + "6666" + "</b>.</div>\n<div>author:Loveyless.</div>") //邮件内容

  // err := e.Send("smtp.163.com:465", smtp.PlainAuth("", "githubbyloveyless@163.com", "NWNXHEUGDOYJJSYG", "smtp.163.com"))
  // e.Send用不了EOF 只能使用SendWithTLS关闭SSL连接就可以了 还有一个函数我也发送失败->SendWithTLS通过可选的tls配置发送电子邮件。

  err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "githubbyloveyless@163.com", "NWNXHEUGDOYJJSYG", "smtp.163.com"), &tls.Config{
    //跳过验证
    InsecureSkipVerify: true,
    ServerName:         "smtp.163.com",
  })
  if err != nil {
    t.Fatal(err)
  }
}
```



### 邮箱验证注册

time.Duration的使用

https://blog.csdn.net/taoshihan/article/details/125924020?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-0-125924020-blog-82706022.t0_edu_mix&spm=1001.2101.3001.4242.1&utm_relevant_index=1



redis存一个 key为邮箱30秒过期 和 value验证码

到时候拿出来比对

```go
set验证码就是Set                                                                这里拿到全局config里面的过期时间 用time.Duration处理一下
err = service.Rdb.Set(ctx, config.RegisterPrefix+emailInfo.Email, emailCode, time.Second*time.Duration(config.RegisterLowTime)).Err()


验证就是Getredis
emailCode, err := service.Rdb.Get(context.Background(), config.RegisterPrefix+info.Email).Result()
```





## websocket服务





官方包里有示例

![](https://cdn.jsdelivr.net/gh/Loveyless/img-clouding/img/20220926010818.png)



### **echo简易版**

比官方的还简易

![](https://cdn.jsdelivr.net/gh/Loveyless/img-clouding/img/20220926014024.png)

```Go
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

//这个里面先不用管
func echo(w http.ResponseWriter, r *http.Request) {

  c, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Print("upgrade:", err)
    return
  }
  defer c.Close()
  for {
    mt, message, err := c.ReadMessage()
    if err != nil {
      log.Println("read:", err)
      break
    }
    log.Printf("recv: %s", message)
    err = c.WriteMessage(mt, message)
    if err != nil {
      log.Println("write:", err)
      break
    }
  }
}


func TestWebsocket(t *testing.T) {
  // /echo接口
  http.HandleFunc("/echo", echo)
  log.Fatal(http.ListenAndServe(*addr, nil))
}
```





### 切片存句柄 一对多

但是会给自己也发一条

```Go
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

//一对多 定义一个wsList切片
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

```





### gin中搭建websocket

```Go
func TestGinWebsocket(t *testing.T) {
  r := gin.Default()
  //websocket
  r.GET("/echo", func(c *gin.Context) {
    //c.Writer实现了http.ResponseWriter接口
    //c.Request实现了*http.Request接口
    echo(c.Writer, c.Request)
  })
  r.Run()
}


//
var addr = flag.String("addr", "localhost:8080", "http service address")

// Upgrader指定升级HTTP连接为WebSocket连接。
// 并发调用Upgrader的方法是安全的。
var upgrader = websocket.Upgrader{}

//一对多 定义一个wsList切片
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

```





### websocket核心 发送/接收 消息

**简易版**

**router**

```Go

  //用户相关的分组 需要验证token
  user := r.Group("/user", MyJwt.FilterToken())
  
  // 发送接收消息
  user.GET("/websocket/message", controller.WebsocketMessage)
```

**websocket controller**

核心就是发送接收消息 

后期在协议前 存入了用户句柄 判读用户是否在带过来的房间 遍历出在线用户的 conn 在write发送消息

```Go
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

```



### websocket遇到的问题

  //自己从请求头解析 发现可以解析 但是下面接收消息的时候 会报错 websocket: close 1006 (abnormal closure): unexpected EOF
  //报错会导致 连接后秒断



  //自己从参数解析 一切正常！

```Go
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

  //自己从请求头解析 发现可以解析 但是下面接收消息的时候 会报错 websocket: close 1006 (abnormal closure): unexpected EOF
  //报错会导致 连接后秒断
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

  //自己从参数解析 一切正常！
  // token := c.Query("tt")
  // if token == "" {
  //   c.JSON(200, gin.H{
  //     "status":  400,
  //     "message": "token不能为空",
  //   })
  //   c.Abort()
  //   return
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
    //   log.Printf("user_identity:%v,room_identity:%v error:%v", claims.Identity, messaageInfo.RoomIdentity, err.Error())
    //   return
    // }
    // // 2.获取在特定房间的在线用户
    // //   查询房间的所有用户
    // userRoomList, err := service.GetUserRoomByRoomIdentity(messaageInfo.RoomIdentity)
    // if err != nil {
    //   log.Printf("room_identity:%v error:%v", messaageInfo.RoomIdentity, err.Error())
    //   return
    // }

    //  遍历出特定房间的在线用户
    // for _, userRoom := range userRoomList {
    //   // 如果房间中所有用户有在线的用户那么就发送消息 注意这里巧妙的写法!!
    //   if conn, ok := wsMap[userRoom.UserIdentity]; ok {
    //     err := conn.WriteMessage(websocket.TextMessage, []byte(messaageInfo.Message))
    //     if err != nil {
    //       log.Println("发送数据失败", err.Error())
    //       return
    //     }
    //   }
    // }
    // 3.保存消息
    // messageBasic := &service.MessageBasic{
    //   UserIdentity: claims.Identity,
    //   RoomIdentity: messaageInfo.RoomIdentity,
    //   Data:         messaageInfo.Message,
    //   CratedAt:     time.Now().Unix(),
    //   UpdatedAt:    time.Now().Unix(),
    // }
    // err = service.InsertOneMessageBasic(messageBasic)
    // if err != nil {
    //   log.Printf("保存消息错误%v", err.Error())
    //   return
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
```





## 查询别人个人资料

查询别人的资料并且是否为好友

get请求/query/detail

还是比较小复杂的 对比两个人有没有相同的房间





## 添加好友/删除好友



1. 添加好友

初步想法

查看是否为好友

不是的话 给两个人弄用一个房间id存在room_basic表中

























