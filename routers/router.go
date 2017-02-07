package routers

import (
	"github.com/astaxie/beego"
	"start/controllers/admin"
	"start/controllers/web"
)

func init() {

	beego.Router("/login", &web.LoginController{})
	beego.Router("/login/check", &web.LoginController{}, "post:Check")
	// 后台url
	beego.Router("/admin", &admin.HomeController{})
	beego.Router("/admin/add", &admin.HomeController{}, "post:Add")
	beego.Router("/admin/release", &admin.HomeController{}, "post:Release")
	beego.Router("/admin/user", &admin.UserController{})
	beego.Router("/admin/sys", &admin.SysController{})
	beego.Router("/admin/upload", &admin.UploadController{})

}
