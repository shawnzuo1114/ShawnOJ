package midware

import (
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func SendVerificationEmail(email string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(8999) + 1000
	code := strconv.Itoa(num)
	host := "smtp.163.com" // 使用163邮箱发送
	port := 465
	username := "ShawnOJ@163.com"
	to := email
	password := "VRa8M94TMCAaYBfu"
	subject := "ShawnOJ验证码"
	body := "您的验证码为：<br><h1>" + code + "</h1><br>五分钟该验证码有效。<br>ShawnOJ"
	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, username, password)
	if err := d.DialAndSend(m); err != nil {
		log.Fatalln(err)
	}
	return code
}
