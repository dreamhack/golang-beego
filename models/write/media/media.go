package media

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbWrite, err = sql.Open("mysql", "root:root@/start?charset=utf8")
)

func UploadMedia(uploadUid string, uploadTime string, fileName string, fileSize int64, isDel int, extName string) (int64, error) {

	stmt, err := dbWrite.Prepare("INSERT FileInfo SET UploadUid=?,UploadTime=?,FileName=?,FileSize=?,IsDel=?,Extname=?")
	res, err := stmt.Exec(uploadUid, uploadTime, fileName, fileSize, isDel, extName)
	checkErr(err)
	lastId, err := res.LastInsertId()
	// dbWrite.Close()

	return lastId, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
