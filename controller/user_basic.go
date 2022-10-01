package controller

import (
	"context"
	"fmt"
	"gin_websocket_test/MyJwt"
	"gin_websocket_test/MyUtils"
	"gin_websocket_test/config"
	"gin_websocket_test/email"
	"gin_websocket_test/service"
	"gin_websocket_test/validator"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v4"
)

//redis需要用到
var ctx = context.Background()

//登录
func Login(c *gin.Context) {

	//获取数据并验证
	userInfo := struct {
		Username string `json:"username" binding:"required,min=3,max=20"`
		Password string `json:"password" binding:"required,min=3,max=20"`
	}{}
	err := c.ShouldBind(&userInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 400,
			"data":   validator.Translate(err),
		})
		c.Abort()
		return
	}

	//查询用户
	u, err := service.GetUserBasicByUsernamePassword(userInfo.Username, userInfo.Password)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "查询不到值",
			"data":    validator.Translate(err),
		})
		c.Abort()
		return
	}

	//生成token
	claims := MyJwt.MyCustomClaims{
		Identity: u.Identity,
		Username: u.Username,
		Email:    u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()), //签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), //生效时间
		},
	}
	token, err := MyJwt.MakeToken(claims)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  500,
			"message": "jwt内部错误" + err.Error(),
			"data":    nil,
		})
		c.Abort()
		return
	}

	//无错返回
	c.JSON(200, gin.H{
		"status":  200,
		"message": "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

//用户详情 根据token负载
func UserDetail(c *gin.Context) {

	// //获取token
	// tokenString := c.GetHeader("token")

	// //解析负载
	// userInfo, err := MyJwt.ParseToken(tokenString)
	// if err != nil {
	// 	c.JSON(200, gin.H{
	// 		"status":  400,
	// 		"message": "查询失败",
	// 		"data":    validator.Translate(err),
	// 	})
	// 	c.Abort()
	// 	return
	// }

	// 直接获取负载不采用上面的获取解析(在中间件已经解析过了)
	userInfo := c.MustGet("claims").(*MyJwt.MyCustomClaims) //需要断言 还可以使用get

	//通过负载中的id查询
	// id string转objectid 通过id查询 但是mongo需要的id不是字符串类型 转一下 !!在websocket后 这里不需要转了 因为不用mongodb的id了 用自己的字段Identity
	// id, err := primitive.ObjectIDFromHex(userInfo.Id)
	// if err != nil {
	// 	c.JSON(200, gin.H{
	// 		"status":  400,
	// 		"message": "查询失败",
	// 		"data":    err.Error(),
	// 	})
	// 	c.Abort()
	// 	return
	// }

	//查询
	ub, err := service.GetUserBasicByIdentity(userInfo.Identity)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "查询失败",
			"data":    validator.Translate(err),
		})
		c.Abort()
		return
	}

	//返回
	c.JSON(200, gin.H{
		"status":  200,
		"message": "",
		"data":    ub,
	})
}

//用户详情 根据query参数
func UserQueryDetail(c *gin.Context) {

	claims := c.MustGet("claims").(*MyJwt.MyCustomClaims) //需要断言 还可以使用get

	username := c.Query("user_name")
	if username == "" {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "参数不正确",
			"data":    gin.H{},
		})
		c.Abort()
		return
	}

	//根据用户名查询对方info
	ub, err := service.GetUserBasicByUsernameInfo(username)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "查询失败",
			"data":    err.Error(),
		})
		c.Abort()
		return
	}
	fmt.Println(ub)
	//过滤密码等信息
	friendInfo := new(struct {
		Nicname  string `json:"nicname" bson:"nicname"`
		Sex      int    `json:"sex" bson:"sex"`
		Email    string `json:"email" bson:"email"`
		Avatar   string `json:"avatar" bson:"avatar"`
		IsFriend bool   `json:"is_friend" bson:"is_friend"` //是否是好友 true是 false否
	})
	friendInfo.Nicname = ub.Nicname
	friendInfo.Sex = ub.Sex
	friendInfo.Email = ub.Email
	friendInfo.Avatar = ub.Avatar
	//查询是否好友
	b, _ := service.JudgeUserIsFriend(claims.Identity, ub.Identity)
	friendInfo.IsFriend = b

	//返回
	c.JSON(200, gin.H{
		"status":  200,
		"message": "查询成功",
		"data":    friendInfo,
	})

}

//注册
func Register(c *gin.Context) {
	//匿名结构体
	info := struct {
		Username string `json:"username" binding:"required,min=3,max=20"`
		Password string `json:"password" binding:"required,min=3,max=20"`
		Email    string `json:"email" binding:"required,email"`
		Code     string `json:"code" binding:"required"`
	}{}
	err := c.ShouldBindWith(&info, binding.JSON)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "参数错误",
			"data":    validator.Translate(err),
		})
		c.Abort()
		return
	}

	//验证验证码 先验证验证码 可以减压数据库
	emailCode, err := service.Rdb.Get(context.Background(), config.RegisterPrefix+info.Email).Result()
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "验证码已过期",
			"data":    err.Error(),
		})
		c.Abort()
		return
	}
	if emailCode != info.Code {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "验证码错误",
			"data":    gin.H{},
		})
		c.Abort()
		return
	}

	// 判断账号是否唯一
	_, err = service.GetUserBasicByUsername(info.Username)
	exist := service.GetUserBasicByEmail(info.Email)
	if err == nil || exist {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "用户名或邮箱已存在",
			"data":    gin.H{},
		})
		c.Abort()
		return
	}

	//存入数据库
	userBasic := service.UserBasic{
		Identity:  MyUtils.Uuid(), //id随机
		Username:  info.Username,
		Password:  MyUtils.Md5(info.Password), //密码做md5处理
		Email:     info.Email,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	err = service.InsertUserBasic(&userBasic)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "注册失败",
			"data":    err,
		})
		c.Abort()
		return
	}

	//无错返回
	c.JSON(200, gin.H{
		"status":  200,
		"message": "注册成功!",
		"data": gin.H{
			"username": info.Username,
		},
	})

}

//发送验证码
func SendCode(c *gin.Context) {

	//获取请求数据并验证
	emailInfo := struct {
		Email string `json:"email" binding:"required,email"`
	}{}
	err := c.ShouldBind(&emailInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 400,
			"data":   validator.Translate(err),
		})
		c.Abort()
		return
	}

	//查询是否存在
	exist := service.GetUserBasicByEmail(emailInfo.Email)
	if exist {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "邮箱已存在",
			"data":    gin.H{},
		})
		c.Abort()
		return
	}

	//发送验证码
	//生成
	emailCode := email.GetEmailCode()
	//发送
	ok := email.SendEmailCode(emailInfo.Email, emailCode)
	if !ok {
		log.Println("发送验证码失败", err)
		c.JSON(200, gin.H{
			"status":  400,
			"message": "发送失败",
			"data":    err.Error(),
		})
		c.Abort()
		return
	}
	//存入redis 30秒过期
	err = service.Rdb.Set(ctx, config.RegisterPrefix+emailInfo.Email, emailCode, time.Second*time.Duration(config.RegisterLowTime)).Err()
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "redis内部出错",
			"data":    err.Error(),
		})
		log.Println("redis内部出错", err)
		c.Abort()
		return
	}

	//返回
	c.JSON(200, gin.H{
		"status":  200,
		"message": "发送成功",
		"data":    nil,
	})
}

//添加用户
func UserAdd(c *gin.Context) {

	//获取请求数据并验证
	info := new(struct {
		Username string `json:"user_name" binding:"required,min=3,max=20"`
	})
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "请求失败",
			"data":    validator.Translate(err),
		})
		c.Abort()
		return
	}

	//判断账号是否存在
	ub, err := service.GetUserBasicByUsername(info.Username)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "用户不存在",
			"data":    gin.H{},
		})
		c.Abort()
		return
	}

	//是否好友关系
	claims := c.MustGet("claims").(*MyJwt.MyCustomClaims)
	b, _ := service.JudgeUserIsFriend(claims.Identity, ub.Identity)
	if b {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "已是好友",
			"data":    gin.H{},
		})
		c.Abort()
		return
	}

	//保存房间记录  这里就是加了一条房间信息 创建者就是加好友的发起者
	rb := &service.RoomBasic{
		// Number: ,房间号
		// Name: ,  房间名
		// Info: ,  房间简介    单聊没有这三个
		Identity:     MyUtils.Uuid(),  //房间唯一标识
		UserIdentity: claims.Identity, //房间创建者id 也就是我的id (这里我的id就是token里的identity)
		CratedAt:     time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}
	err = service.InsertOneRoomBasic(rb)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "添加room_basic失败",
			"data":    err,
		})
		c.Abort()
		return
	}

	//保存用户与房间的关联记录 这里就是拿上面的房间id 与我的id/好友id 存两条记录
	//第一个
	ur := &service.UserRoom{
		UserIdentity: claims.Identity, //我的id
		RoomIdentity: rb.Identity,     //房间id
		RoomType:     1,               //单聊
		CratedAt:     time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}
	err = service.InsertOneUserRoom(ur)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "添加user_room失败",
			"data":    err,
		})
		c.Abort()
		return
	}
	//第二个
	ur = &service.UserRoom{
		UserIdentity: ub.Identity, //对方的id
		RoomIdentity: rb.Identity, //房间id
		RoomType:     1,           //单聊
		CratedAt:     time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}
	err = service.InsertOneUserRoom(ur)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "添加user_room失败",
			"data":    err,
		})
		c.Abort()
		return
	}

	//返回
	c.JSON(200, gin.H{
		"status":  200,
		"message": "添加成功",
		"data":    gin.H{},
	})
}

//删除用户
func UserDelete(c *gin.Context) {
	claim := c.MustGet("claims").(*MyJwt.MyCustomClaims)

	ubId := c.Query("user_identity")
	if ubId == "" {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "无此用户",
			"data":    gin.H{},
		})
		c.Abort()
		return
	}

	//获取房间identity 通过双方id
	fmt.Println(claim.Identity, ubId)
	userRoomIdentity := service.GetUserRoomIdentity(claim.Identity, ubId)
	if userRoomIdentity == "" {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "与对方不是好友",
			"data":    gin.H{},
		})
		c.Abort()
		return
	}

	//通过房间id删除俩人的关联记录 RoomBasic
	err := service.DeleteRoomBasicByRoomIdentity(userRoomIdentity)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "房间不存在",
			"data":    err,
		})
		c.Abort()
		return
	}

	//通过房间id和双方id 删除俩人的关联记录 UserRoom
	err = service.DeleteUserRoomByRoomIdentity(userRoomIdentity)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "房间不存在",
			"data":    err,
		})
		c.Abort()
		return
	}

	//返回
	c.JSON(200, gin.H{
		"status":  200,
		"message": "删除成功",
		"data":    gin.H{},
	})

}
