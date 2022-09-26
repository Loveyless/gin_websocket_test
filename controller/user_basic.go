package controller

import (
	"gin_websocket_test/MyJwt"
	"gin_websocket_test/email"
	"gin_websocket_test/service"
	"gin_websocket_test/validator"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

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
	ok := email.SendEmailCode(emailInfo.Email, "123456")
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

	//返回
	c.JSON(200, gin.H{
		"status":  200,
		"message": "发送成功",
		"data":    nil,
	})
}
