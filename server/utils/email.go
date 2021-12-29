package utils

import (
	"github.com/golang/glog"
	"gopkg.in/gomail.v2"
)

type GMail struct {
}

func NewGMail() *GMail {
	return &GMail{}
}

func SendQQ(email, title, content string) {
	fun := "GMail.SendQQ-->"
	m := gomail.NewMessage()
	m.SetHeader("From", "1208579668@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", content)

	d := gomail.NewDialer("smtp.qq.com", 465, "1208579668@qq.com", "kqiufdxbnnbligdg")
	if err := d.DialAndSend(m); err != nil {
		glog.Errorf("%s send to %s err: %v", fun, email, err)
		return
		// TODO 重试
	}
}
