package svn

import (
	"bytes"
	"os/exec"
)

const (
	CodePath = "/static/code/"
)

func CheckOut(getUrl string, getName string, userName string, passWord string) (string, error) {

	var err error

	cmd := exec.Command("svn checkout", getUrl, CodePath+getName, userName, passWord)
	w := bytes.NewBuffer(nil)
	cmd.Stderr = w
	err = cmd.Run()
	svnOutput := string(w.Bytes())

	return svnOutput, err
}
