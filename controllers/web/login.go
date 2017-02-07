package web

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

var cpt *captcha.Captcha

func init() {

	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {

	this.TplNames = "login.tpl"

}

func (this *LoginController) Check() {

	jsoninfo := this.GetString("name")
	city := this.GetString("city")

	this.Ctx.WriteString(jsoninfo)
	this.Ctx.WriteString(city)
	this.Ctx.WriteString("你好")
	// this.Data["json"] = jsoninfo
	// this.Data["json"] = city
	// this.ServeJson()
	// this.StopRun()
	this.TplNames = "login.tpl"
	// this.Data["Success"] = cpt.VerifyReq(this.Ctx.Request)
}
