package main
 
import (
	"fmt"
	"time"
)
 
func main() {
    //时间戳
    t := time.Now().Unix()
    fmt.Println(t)
     
    //时间戳到具体显示的转化
    fmt.Println(time.Unix(t, 0).String())
     
    //带纳秒的时间戳
    t = time.Now().UnixNano()
    fmt.Println(t)
    fmt.Println("------------------")
     
    //基本格式化的时间表示
    fmt.Println(time.Now().String())
     
    fmt.Println(time.Now().Format("2006year 01month 02day"))
 
}