package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"start/extension/bash/svn"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {

	this.Layout = "admin/layout.html"
	this.TplNames = "admin/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["SideBar"] = "common/sidebar.html"
	this.LayoutSections["Scripts"] = "admin/scripts.html"

}

func (this *HomeController) Add() {

	get_url := this.GetString("url")
	get_proname := this.GetString("proname")
	get_username := this.GetString("username")
	get_password := this.GetString("password")

	// this.Ctx.WriteString(getUrl)
	// this.StopRun()

	svnOutput, err := svn.CheckOut(get_url, get_proname, get_username, get_password)
	svnLog := logs.NewLogger(1000)
	svnLog.SetLogger("file", `{"filename":"logs/svn.log"}`)
	svnLog.Info("#TASK:" + "FILEID:")
	svnLog.Debug(svnOutput)
	svnLog.Error(err.Error() + "\n" + "##############################################")
}

func (this *HomeController) Release() {

	this.Layout = "admin/layout.html"
	this.TplNames = "admin/release.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["SideBar"] = "common/sidebar.html"
	this.LayoutSections["Scripts"] = "admin/scripts.html"

}
