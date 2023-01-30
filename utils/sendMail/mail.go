package sendMail

import (
	"gopartsrv/public/config"
	"gopkg.in/gomail.v2"
)

func testSend() {
	email := map[string]interface{}{
		"html":   VerificationCodeHtml("旅客", "123456"),
		"email":  "2461501270@qq.com",
		"object": "验证码",
	}
	SendMail(email)
}

// 发送邮件
func SendMail(email map[string]interface{}) error {
	list := config.GetEmailConfig()["jerome"].(map[string]interface{})
	m := gomail.NewMessage()
	m.SetHeader("From", list["username"].(string))
	m.SetHeader("To", email["email"].(string))
	m.SetHeader("Subject", email["object"].(string)) //设置邮件主题
	m.SetBody("text/html", email["html"].(string))   //设置邮件正文
	// 第一个参数是host 第三个参数是发送邮箱，第四个参数 是邮箱密码
	d := gomail.NewDialer(list["host"].(string), int(list["port"].(int64)), list["username"].(string), list["password"].(string))
	if err := d.DialAndSend(m); err != nil {
		panic(err)
		return err
	}
	return nil
}
