package main

import (
	"net/smtp"
	//"github.com/smtc/glog"
	"fmt"
)

func main() {
	auth := smtp.PlainAuth("", "", "", "smtp.qiye.163.com")
	to := []string{"", ""}
	msg := []byte("To: \r\n" +
		"Subject: 测试golang邮件发送 \r\n" +
		"Content-Type: text/plain; charset=UTF-8" + "\r\n\r\n")
	//for _, value := range sendMailContent {
	msg = append(msg, []byte("测试邮件发送")...)
	//}
	err := smtp.SendMail("smtp.qiye.163.com:25", auth, "", to, msg)
	if err != nil {
		//glog.Error("邮件发送失败! err: %s \n", err.Error())
		fmt.Println(err.Error())
		return
	}
	fmt.Println("successfully")
}
