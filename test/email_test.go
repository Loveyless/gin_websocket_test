package test

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"testing"
	"time"

	"github.com/jordan-wright/email"
)

//NWNXHEUGDOYJJSYG
func TestSendEmail(t *testing.T) {
	e := email.NewEmail()                                                                                                                                      //创建一个新的邮件
	e.From = "Loveyless <githubbyloveyless@163.com>"                                                                                                           // 发件人
	e.To = []string{"709251884@qq.com"}                                                                                                                        // 收件人 可以多个
	e.Subject = "验证码已发送，请查收"                                                                                                                                   //邮件主题
	e.HTML = []byte("<div>Your verification code:<b>" + "666" + "</b>.</div>\n<div>author:Loveyless.</div>\n<a src='https://github.com/Loveyless'>github</a>") //邮件内容

	// err := e.Send("smtp.163.com:465", smtp.PlainAuth("", "githubbyloveyless@163.com", "NWNXHEUGDOYJJSYG", "smtp.163.com"))
	// e.Send用不了EOF 只能使用SendWithTLS关闭SSL连接就可以了 还有一个函数我也发送失败->SendWithTLS通过可选的tls配置发送电子邮件。

	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "githubbyloveyless@163.com", "NWNXHEUGDOYJJSYG", "smtp.163.com"), &tls.Config{
		//跳过验证
		InsecureSkipVerify: true,
		ServerName:         "smtp.163.com",
	})
	if err != nil {
		t.Fatal(err)
	}
}

// https://github.com/jordan-wright/email
func TestEmail(t *testing.T) {
	e := email.NewEmail()                                     //创建一个新的邮件
	e.From = "Jordan Wright <test@gmail.com>"                 // 发件人
	e.To = []string{"test@example.com"}                       // 收件人 可以多个
	e.Bcc = []string{"test_bcc@example.com"}                  // Bcc: Blind Carbon Copy
	e.Cc = []string{"test_cc@example.com"}                    // Cc: Carbon Copy
	e.Subject = "Awesome Subject"                             //邮件主题
	e.Text = []byte("Text Body is, of course, supported!")    //邮件内容
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>") //邮件内容
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "GithubByLoveyless@163.com", "password123", "smtp.163.com"))
}

//生成验证码
func TestGetEmailCode(t *testing.T) {
	//种子
	rand.Seed(time.Now().UnixNano())
	res := ""
	//长度为6
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	fmt.Println(res)
}
