# websocket聊天室

## 前置准备
```
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/websocket
go get go.mongodb.org/mongo-driver/mongo
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
|-MyJwt       token中间件
|-router      路由 分组/拦截
|-service     操作数据的函数
    mongodb.go    连接mongo数据库
|-utils       工具库
|-validator   将gin验证器错误翻译成中文
|-main.go
```