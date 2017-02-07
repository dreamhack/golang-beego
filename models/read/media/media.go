package media

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbRead, err = sql.Open("mysql", "root:root@/start?charset=utf8")
)

func StartQueue() (StartId uint64, Filename string, Extname string, FileSize int64, err error) {

	checkErr(err)
	rows, err := dbRead.Query("SELECT FID,FileName,Extname,FileSize FROM FileInfo where FID=(SELECT MIN(FID) FROM FileInfo WHERE IsTranscode = 0);")

	for rows.Next() {

		err = rows.Scan(&StartId, &Filename, &Extname, &FileSize)
	}
	dbRead.Close()
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
