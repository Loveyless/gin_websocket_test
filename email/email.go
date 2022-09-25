package email

import (
	"crypto/tls"
	"log"
	"net/smtp"

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
	e.HTML = []byte("<div>Your verification code:<b>" + code + "</b>.</div>\n<div>author:Loveyless.</div>") //邮件内容
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "githubbyloveyless@163.com", "NWNXHEUGDOYJJSYG", "smtp.163.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		log.Println(err)
	}
	return err == nil
}

// https://github.com/jordan-wright/email
func TestEmail() {
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
