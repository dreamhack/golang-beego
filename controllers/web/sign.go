package web

import (
	"github.com/astaxie/beego"

	// "start/extension/mail"
)

type SignController struct {
	beego.Controller
}

func (this *SignController) Get() {

	// email := mail.NewEmail("dreamhack@live.com", "title", "会资料下载地址")
	// err := mail.SendEmail(email)
	// fmt.Println(err)
	// this.Data["json"] = "ddd,5555"
	// this.ServeJson()
	// this.StopRun()
	this.TplNames = "sign.tpl"
}
