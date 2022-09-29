package email

import (
	"crypto/tls"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/jordan-wright/email"
)

// SendCode
/* 发送验证码
要发送的邮箱 （可以改成切片 目前只支持一个）
验证码
*/
func SendEmailCode(toUserEmail, code string) bool {
	e := email.NewEmail()
	e.From = "Loveyless <githubbyloveyless@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "The verification code has been sent, please check"
	e.HTML = []byte("<div>Your verification code:<b>" + code + "</b>.</div>\n<div>author:Loveyless.</div>\n<a src='https://github.com/Loveyless'>github</a>") //邮件内容
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "githubbyloveyless@163.com", "NWNXHEUGDOYJJSYG", "smtp.163.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		log.Println(err)
	}
	return err == nil
}

//生成验证码
func GetEmailCode() string {
	//种子
	rand.Seed(time.Now().UnixNano())
	res := ""
	//长度为6
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	return res
}
