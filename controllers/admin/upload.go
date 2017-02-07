package admin

import (
	"github.com/astaxie/beego"
	"os"
	"path/filepath"
	"start/models/write/media"
	"time"
)

type UploadController struct {
	beego.Controller
}

type FileInfo struct {
	uploadUid  string
	uploadTime string
	fileName   string
	fileSize   int64
	isDel      int
	extName    string
}

func (this *UploadController) Get() {

	this.TplNames = "common/upload.tpl"
}

func (this *UploadController) Post() {

	nowTime := time.Now().Format("20060102150405000")
	this.EnableRender = false
	if this.Ctx.Input.IsUpload() {

		filename := "123" + nowTime
		mf := this.Ctx.Request.MultipartForm
		extName := filepath.Ext(mf.Value["name"][0])
		saveFilename := "static/media/" + filename + extName
		this.SaveToFile("file", saveFilename)
		setFileInfo := FileInfo{"123", nowTime, filename, FileSize(saveFilename), 0, extName}
		getLatid, err := media.UploadMedia(setFileInfo.uploadUid, setFileInfo.uploadTime, setFileInfo.fileName, setFileInfo.fileSize, setFileInfo.isDel, setFileInfo.extName)
		beego.Trace(getLatid, err)
	}

	this.TplNames = "common/footer.tpl"
	return
}

func FileSize(file string) int64 {
	f, _ := os.Stat(file)
	return f.Size()
}
