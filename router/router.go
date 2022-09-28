package router

import (
	"gin_websocket_test/MyJwt"
	"gin_websocket_test/controller"
	"gin_websocket_test/cors"

	"github.com/gin-gonic/gin"
)

func Start() {

	r := gin.Default()
	r.Use(cors.Cors())

	//登录
	r.POST("/login", controller.Login)
	//发送验证码
	r.POST("/send/code", controller.SendCode)

	//用户相关的分组 需要验证token
	user := r.Group("/user", MyJwt.FilterToken())

	// 用户详情
	user.GET("/detail", controller.UserDetail)
	// 发送接收消息
	user.GET("/websocket/message", controller.WebsocketMessage)

	//用户聊天记录列表
	user.GET("/chat/list", controller.ChatList)

	r.Run()
}
