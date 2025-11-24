package mail

import (
	"ShawnOJ/utils/setting"
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

func SendVerifyCode(email string) (code int, err error) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", setting.Conf.Sender)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "ShawnOJ验证码")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code = r.Intn(900000) + 100000

	emailBody := fmt.Sprintf(`
    <div style="font-family: Arial, sans-serif; font-size: 16px; color: #333;">
        <h2>验证码邮件</h2>
        <p>您好，您的邮箱验证码为：</p>
        <p style="font-size: 24px; font-weight: bold; color: #007BFF;">%d</p>
        <p>验证码有效期 5 分钟，请勿泄露给他人。</p>
    </div>`, code)
	msg.SetBody("text/html", emailBody)

	n := gomail.NewDialer("smtp.163.com", 465, setting.Conf.Sender, setting.Conf.EmailConfig.Password)
	if err = n.DialAndSend(msg); err != nil {
		return 0, err
	}

	return code, nil
}
