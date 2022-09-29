package controller

import (
	"context"
	"gin_websocket_test/MyJwt"
	"gin_websocket_test/config"
	"gin_websocket_test/email"
	"gin_websocket_test/service"
	"gin_websocket_test/validator"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-redis/redis/v9/internal/util"
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

//用户详情
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

//注册
func Register(c *gin.Context) {
	//匿名结构体
	info := struct {
		Username string `json:"username" binding:"required,min=3,max=20"`
		Password string `json:"password" binding:"required,min=3,max=20"`
		Email    string `json:"email" binding:"required,email"`
		Code     string `json:"code" binding:"required"`
	}{}
	err := c.ShouldBindWith(info, binding.JSON)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "参数错误",
			"data":    validator.Translate(err),
		})
		c.Abort()
		return
	}

	//验证验证码
	emailCode, err := service.Rdb.Get(context.Background(), config.RegisterPrefix+info.Email).Result()
	if err != nil {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "验证码已过期",
			"data":    err.Error(),
		})
		c.Abort()
		return
	}
	if emailCode != info.Code {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "验证码错误",
			"data":    gin.H{},
		})
		c.Abort()
		return
	}

	// 判断账号是否唯一
	exist1 := service.GetUserBasicByUsername(info.Username)
	exist2 := service.GetUserBasicByEmail(info.Email)
	if exist1 || exist2 {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "用户名或邮箱已存在",
			"data":    gin.H{},
		})
		c.Abort()
		return
	}

	//存入数据库
	userBasic := service.UserBasic{
		Identity: uuid.NewV4().String(), //生成uuid
		Username: info.Username,
		Password: util.Md5(info.Password),
		Email:    info.Email,
	}

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
