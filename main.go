package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "start/extension/mail"
	_ "start/routers"
)

func main() {

	// sum := 1
	// for sum < 10000 {

	// 	email := mail.NewEmail("157261119@qq.com", "米神你好", "欢迎米神回来")
	// 	sum++

	// 	err := mail.SendEmail(email)
	// 	fmt.Println(err)
	// 	fmt.Println("已经发送: %v封邮件！", sum)

	// }
	beego.SetStaticPath("static/", "static")
	// beego.BeeLogger.DelLogger("console")
	beego.SetLogger("file", `{"filename":"logs/start.log"}`)
	beego.SetLogFuncCall(true)
	orm.Debug = true
	beego.Run()
}
