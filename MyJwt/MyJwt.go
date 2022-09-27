package MyJwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

//传入负载生成token
//func NewJwt  返回token字符串

//验证token错误(过期 格式错误)
//func ParseJwt  返回解析的负载结构体

//过滤路径 //这个可以使用路由分组实现 某个分组都带上这个中间件
// var FilterSlice = []string{"/login"}

//加密key
var key = []byte("loveyless")

//验证负载结构体
type MyCustomClaims struct {
	Identity string `json:"identity"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// Password             string `json:"password"`
	jwt.RegisteredClaims // 注册声明
}

//拦截token中间件
func FilterToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")

		// 过滤不需要的路径 可以使用路由分组实现
		// for _, v := range FilterSlice {
		// 	if c.Request.URL.Path == v {
		// 		c.Next()
		// 		return
		// 	}
		// }

		// 判空
		if token == "" {
			c.JSON(200, gin.H{
				"status": 400,
				"msg":    "token不能为空",
			})
			c.Abort()
			return
		}

		// 验证token是否错误
		_, err := ParseToken(token)
		if err != nil {
			c.JSON(200, gin.H{
				"status": 400,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}

		//可选步骤 解析一下token 把值挂在上下文中 到时候直接Get获取 但是获取时 是any类型 还要转回来 *MyCustomClaims
		claims, _ := ParseToken(token)
		// fmt.Println("解析后的负载", claims)

		// // 挂载到上下文中
		c.Set("claims", claims)

	}
}

/*
 * 生成token
 * claims负载
 *
 * tokenString 返回的token字符串
// 示例
// // Create the claims
// claims := MyCustomClaims{
//   "bar",
//   jwt.RegisteredClaims{
//     // A usual scenario is to set the expiration time relative to the current time
//     ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
//     IssuedAt:  jwt.NewNumericDate(time.Now()),
//     NotBefore: jwt.NewNumericDate(time.Now()),
//     Issuer:    "test",
//     Subject:   "somebody",
//     ID:        "1",
//     Audience:  []string{"somebody_else"},
//   },
// }

//或者
// claims := MyJwt.MyCustomClaims{
//   Id:       value.Id,
//   Username: value.Username,
//   Password: value.Password,
//   RegisteredClaims: jwt.RegisteredClaims{
//     ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
//     IssuedAt:  jwt.NewNumericDate(time.Now()),
//     NotBefore: jwt.NewNumericDate(time.Now()),
//     Issuer:    "test",
//     Subject:   "somebody",
//     ID:        "1",
//     Audience:  []string{"somebody_else"},
//   },
// }
// toekn, _ := MyJwt.NewJwt(claims)

// 创建声明，同时省略一些可选字段
// claims = MyCustomClaims{
//   "bar",
//   jwt.RegisteredClaims{
//     // Also fixed dates can be used for the NumericDate
//     ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
//     Issuer:    "test",
//   },
// }
*/
func MakeToken(claims MyCustomClaims) (string, error) {

	// 携带负载
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 加密
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

/*
 * 解析token
 * tokenString 要解析的token字符串
 * 示例
// token := c.Request.Header.Get("token")
// res, err := MyJwt.ParseJwt(token)
*/
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	// 解析token
	//参数 token 解密的模板 函数 (返回你的key)
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	// 解析错误
	if err != nil {
		return nil, err
	}

	// 解析token
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		// 返回结构体地址
		return claims, nil
	} else {
		return nil, err
	}

}
