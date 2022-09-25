package test

import (
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

//NWNXHEUGDOYJJSYG
func TestSendEmail(t *testing.T) {
	e := email.NewEmail()                                                                                     //创建一个新的邮件
	e.From = "Loveyless <githubbyloveyless@163.com>"                                                          // 发件人
	e.To = []string{"709251884@qq.com"}                                                                       // 收件人 可以多个
	e.Subject = "验证码已发送，请查收"                                                                                  //邮件主题
	e.HTML = []byte("<div>Your verification code:<b>" + "6666" + "</b>.</div>\n<div>author:Loveyless.</div>") //邮件内容

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
