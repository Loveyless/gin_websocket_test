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
    "number":"房间号",
    "name":"房间名",
    "info":"简介",
    "user_identity":"房间创建者的唯一标识",
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
|-controller  基础逻辑
|-cors        跨域中间件
|-email       封装发送邮箱
|-MyJwt       token中间件
|-router      路由 分组/拦截
|-service     操作数据的函数
    mongodb.go    连接mongo数据库
|-test        测试
|-utils       工具库
|-validator   将gin验证器错误翻译成中文
|-main.go
```



## 邮箱库的使用

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