package router

import (
	"gin_websocket_test/MyJwt"
	"gin_websocket_test/MyUtils"
	"gin_websocket_test/controller"
	"gin_websocket_test/cors"

	"github.com/gin-gonic/gin"
)

func Start() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Cors(), MyUtils.CountTiming)

	{
		//登录
		r.POST("/login", controller.Login)
		//注册
		r.POST("/register", controller.Register)
		//发送验证码
		r.POST("/send/code", controller.SendCode)
	}

	//用户相关的分组 需要验证token
	user := r.Group("/user", MyJwt.FilterToken())

	{
		// 用户详情 从token中的identity获取
		user.GET("/detail", controller.UserDetail)
		// 用户详情 用传递的identity获取
		user.GET("/query/detail", controller.UserQueryDetail)
		// 发送接收消息 不在user分组 因为user分组需要验证token js里的websocket不能随便带请求头 只能自己解析
		r.GET("/websocket/message", controller.WebsocketMessage)
		// 群聊 目前用这个
		r.GET("/websocket/all", controller.WebsocketAll)
		//用户聊天记录列表
		user.GET("/chat/list", controller.ChatList)

		//添加用户
		user.POST("/add", controller.UserAdd)
		//删除用户
		user.DELETE("/delete", controller.UserDelete)
	}

	r.Run()
}
