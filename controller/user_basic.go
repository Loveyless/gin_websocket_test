package controller

import (
	"gin_websocket_test/MyJwt"
	"gin_websocket_test/service"
	"gin_websocket_test/validator"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		Id:       u.Id,
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

	//从token中获取负载
	tokenString := c.GetHeader("token")

	//解析负载
	t, err := MyJwt.ParseToken(tokenString)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "查询失败",
			"data":    validator.Translate(err),
		})
		c.Abort()
		return
	}

	//通过负载中的id查询
	//通过id查询 但是mongo需要的id不是字符串类型 转一下
	id, err := primitive.ObjectIDFromHex(t.Id)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": "查询失败",
			"data":    err.Error(),
		})
		c.Abort()
		return
	}

	//查询
	ub, err := service.GetUserBasicByIdentity(id)
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
