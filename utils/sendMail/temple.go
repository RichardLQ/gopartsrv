package sendMail

import "fmt"

//验证码模版
func VerificationCodeHtml(name, code string) string {
	return fmt.Sprintf(`<div>
		<div>
			敬爱的%s，您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您的验证码是%s</p>
			<p>五分钟内有效</p>
		</div>
		<div>
			<p>此为官方邮箱，仅为验证邮箱</p>
		</div>	
	</div>`, name, code)
}
