package controller

import (
	"gin_websocket_test/MyJwt"
	"gin_websocket_test/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

//获取聊天记录
func ChatList(c *gin.Context) {
	// 1. 判断他是否属于这个房间
	// 2. 查出聊天记录

	claims := c.MustGet("claims").(*MyJwt.MyCustomClaims)
	roomIdentity := c.Query("room_identity")
	if roomIdentity == "" {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "房间号不能为空",
			"data":    gin.H{},
		})
		c.Abort()
		return
	}

	//判断用户是否属于这个房间
	_, err := service.GetUserRoomByUserIdentityRoomIdentity(claims.Identity, roomIdentity)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "无效的房间号",
			"data":    err,
		})
		c.Abort()
		return
	}

	//页码
	pageIndex, err1 := strconv.ParseInt(c.Query("page_index"), 10, 32)
	//查询条数
	limit, err2 := strconv.ParseInt(c.Query("limit"), 10, 32)
	if err1 != nil || err2 != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "页码错误",
			"data":    err,
		})
		c.Abort()
		return
	}
	//偏移量
	offSet := (pageIndex - 1) * limit

	//获取聊天记录
	messagebasicList, err := service.GetMessageListByRoomIdentity(roomIdentity, limit, offSet)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "查询失败",
			"data":    err,
		})
		c.Abort()
		return
	}

	//返回数据
	c.JSON(200, gin.H{
		"status":  200,
		"message": "查询成功",
		"data": gin.H{
			"message_list": messagebasicList,
		},
	})

}
